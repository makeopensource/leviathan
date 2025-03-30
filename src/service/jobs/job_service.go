package jobs

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	. "github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/file_manager"
	"github.com/makeopensource/leviathan/service/labs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

type JobService struct {
	db          *gorm.DB
	dockerSrv   *docker.DkService
	queue       *JobQueue
	broadcastCh *models.BroadcastChannel
	labSrv      *labs.LabService
	fileManSrv  *file_manager.FileManagerService
}

func NewJobService(
	db *gorm.DB,
	bc *models.BroadcastChannel,
	dockerService *docker.DkService,
	labService *labs.LabService,
	tmpFileService *file_manager.FileManagerService,
) *JobService {
	srv := &JobService{
		db:          db,
		broadcastCh: bc,
		dockerSrv:   dockerService,
		queue:       NewJobQueue(uint(ConcurrentJobs.GetUint64()), db, dockerService),
		labSrv:      labService,
		fileManSrv:  tmpFileService,
	}
	srv.cleanupOrphanJobs()
	return srv
}

func (job *JobService) NewJob(newJob *models.Job, submissionFolderId string) (string, error) {
	if newJob.LabData == nil {
		labData, err := job.getLab(newJob.LabID)
		if err != nil {
			return "", err
		}
		newJob.LabData = labData
	}

	jobId, err := uuid.NewUUID()
	if err != nil {
		return "", ErrLog("failed to generate job ID", err, log.Error())
	}
	newJob.JobId = jobId.String()

	mId := job.dockerSrv.ClientManager.GetLeastJobCountMachineId()

	// job context, so that it can be cancelled, and store sub logger
	ctx := job.queue.NewJobContext(newJob.JobId)

	jobDir, err := CreateTmpJobDir(newJob.JobId, SubmissionFolder.GetStr())
	if err != nil {
		return "", ErrLog("failed to create job dir", err, jog(ctx).Error())
	}

	submissionFolder, err := job.fileManSrv.GetSubmissionPath(submissionFolderId)
	if err != nil {
		return "", err
	}
	defer job.fileManSrv.DeleteFolder(submissionFolderId)

	if err = HardLinkFolder(submissionFolder, jobDir); err != nil {
		return "", ErrLog("unable to copy files to job dir", err, log.Error())
	}
	if err = HardLinkFolder(newJob.LabData.JobFilesDirPath, jobDir); err != nil {
		return "", ErrLog("unable to copy files to job dir", err, log.Error())
	}

	logPath, err, reason := setupLogFile(newJob.JobId)
	if err != nil {
		return "", ErrLog("failed to setup log file: "+reason, err, jog(ctx).Error().Str("reason", reason))
	}

	// setup job metadata
	newJob.MachineId = mId
	newJob.Status = models.Queued
	newJob.OutputLogFilePath = logPath
	newJob.TmpJobFolderPath = jobDir
	newJob.JobCtx = ctx
	newJob.LabData.VerifyJobLimits()
	jog(newJob.JobCtx).Debug().Any("limits", newJob.LabData.JobLimits).Msg("job limits")

	res := job.db.Create(newJob)
	if res.Error != nil {
		return "", ErrLog("failed to save job to db", res.Error, jog(ctx).Error())
	}

	err = job.queue.AddJob(newJob)
	if err != nil {
		return "", err
	}

	return jobId.String(), nil
}

func (job *JobService) CancelJob(jobUuid string) {
	job.queue.CancelJob(jobUuid)
}

// WaitForJobAndLogs blocks until job ends
func (job *JobService) WaitForJobAndLogs(jobUuid string) (*models.Job, string, error) {
	var jobInfo *models.Job
	var outerLogs string

	err := job.StreamJobAndLogs(
		context.Background(),
		jobUuid,
		func(jobInf *models.Job, logs string) error {
			jobInfo = jobInf
			outerLogs = logs
			return nil
		},
	)

	return jobInfo, outerLogs, err
}

func (job *JobService) StreamJobAndLogs(
	ctx context.Context,
	jobUuid string,
	streamFunc func(jobInf *models.Job, logs string) error,
) error {
	jobInfo, complete, flogs, err := job.checkJob(jobUuid)
	if err != nil {
		return err
	}

	// send initial job data
	if jobInfo != nil {
		if err := streamFunc(jobInfo, flogs); err != nil {
			return err
		}
	}
	if complete {
		return nil
	}

	jobInfoCh := job.SubToJob(jobInfo.JobId)
	defer job.UnsubToJob(jobUuid)

	// since job.ListenToJobLogs reads logs in an infinite loop,
	// always cancel this context, exiting the loop
	logContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	logsCh := job.ListenToJobLogs(logContext, jobInfo)

	var logs string
	var jobOk = false

	// Keep listening until channel closes, indicating job is complete
	for {
		if jobOk {
			cancel()                                          // stop updating log channel
			content := ReadLogFile(jobInfo.OutputLogFilePath) // final read from the log file
			if err := streamFunc(jobInfo, content); err != nil {
				return err
			}
			return nil
		}
		select {
		case logsTmp, ok := <-logsCh:
			if ok {
				log.Debug().Msg("job logs changed")
				logs = logsTmp
				if err := streamFunc(jobInfo, logsTmp); err != nil {
					return err
				}
			}
		case jobTmp, ok := <-jobInfoCh:
			jobOk = !ok
			if ok && jobInfo != nil && jobTmp.Status != jobInfo.Status {
				log.Debug().Msgf("job status changed from %s to %s", jobInfo.Status, jobTmp.Status)
				jobInfo = jobTmp
				if err := streamFunc(jobInfo, logs); err != nil {
					return err
				}
			}
		case <-ctx.Done():
			return nil
		}
	}
}

// GetJobStatusAndLogs gets the status once whatever it may be and current logs
func (job *JobService) GetJobStatusAndLogs(jobUuid string) (*models.Job, string, error) {
	jobInfo, _, logs, err := job.checkJob(jobUuid)
	if err != nil {
		return nil, "", err
	}
	return jobInfo, logs, nil
}

func (job *JobService) ListenToJobLogs(ctx context.Context, jobInfo *models.Job) chan string {
	logChannel := make(chan string, 2)
	go func(ctx context.Context) {
		prevLength := 0
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		// keep reading until ctx is done
		for {
			select {
			case <-ticker.C:
				content := ReadLogFile(jobInfo.OutputLogFilePath)
				contLen := len(content)
				if contLen > prevLength { // send if content changed
					log.Debug().Str(JobLogKey, jobInfo.JobId).Msgf("sending log, length changed from %d to %d", prevLength, contLen)
					prevLength = contLen
					logChannel <- content
				}
			case <-ctx.Done():
				close(logChannel)
				log.Debug().Str(JobLogKey, jobInfo.JobId).Msg("stopping listening for logs")
				return
			}
		}
	}(ctx)

	return logChannel
}

func (job *JobService) SubToJob(jobUuid string) chan *models.Job {
	return job.broadcastCh.Subscribe(jobUuid)
}

func (job *JobService) UnsubToJob(jobUuid string) {
	job.broadcastCh.Unsubscribe(jobUuid)
}

func (job *JobService) checkJob(jobUuid string) (*models.Job, bool, string, error) {
	jobInf, err := job.getJobFromDB(jobUuid)
	if err != nil {
		return nil, false, "", fmt.Errorf("failed to get job info")
	}

	if jobInf.Status.Done() {
		log.Debug().Str(JobLogKey, jobUuid).Msg("job is already done")

		content := ReadLogFile(jobInf.OutputLogFilePath)
		return jobInf, true, content, nil
	} else {
		return jobInf, false, "", nil
	}
}

func (job *JobService) getJobFromDB(jobUuid string) (*models.Job, error) {
	var jobInfo *models.Job
	res := job.db.First(&jobInfo, "job_id = ?", jobUuid)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get job info from db")
	}
	return jobInfo, nil
}

// removes any job left in an 'active' state before application start,
// fail any jobs that were running before leviathan was able to process them (for whatever reason)
//
// for example machine running leviathan shutdown unexpectedly or leviathan had an unrecoverable error
func (job *JobService) cleanupOrphanJobs() {
	var orphanJobs []*models.Job
	res := job.db.Preload("LabData").
		Where("status = ?", string(models.Queued)).
		Or("status = ?", string(models.Running)).
		Or("status = ?", string(models.Preparing)).
		Find(&orphanJobs)
	if res.Error != nil {
		log.Warn().Err(res.Error).Msgf("Failed to query database for orphan jobs")
		return
	}

	for _, orphan := range orphanJobs {
		if orphan.ContainerId != "" {
			client, err := job.dockerSrv.ClientManager.GetClientById(orphan.MachineId)
			if err != nil {
				log.Warn().Err(err).
					Msgf("unable to find machine: %s ,job: %s was running on", orphan.MachineId, orphan.JobId)
				continue
			}
			err = client.RemoveContainer(orphan.ContainerId, true, true)
			if err != nil {
				log.Warn().Err(err).Str("containerID", orphan.ContainerId).Msg("unable to remove orphan container")
			}
		}

		dkFp := orphan.LabData.DockerFilePath
		if dkFp != "" {
			err := os.RemoveAll(dkFp)
			if err != nil {
				log.Warn().Err(err).Str("dir", dkFp).Msg("unable to remove orphan dockerfile")
			}
		}

		if orphan.TmpJobFolderPath != "" {
			tmpFold := filepath.Dir(orphan.TmpJobFolderPath) // get the dir above autolab subdir
			err := os.RemoveAll(tmpFold)
			if err != nil {
				log.Warn().Err(err).Str("dir", tmpFold).Msg("unable to remove orphan tmp job directory")
			}
		}

		orphan.Status = models.Failed
		orphan.StatusMessage = "job was unable to be processed due to an internal server error"
		res = job.db.Save(orphan)
		if res.Error != nil {
			log.Warn().Err(res.Error).Msg("unable to update orphan job status")
		}
	}

	if len(orphanJobs) != 0 {
		log.Info().Msgf("Cleaned up %d orphan jobs", len(orphanJobs))
	}
}

// setupLogFile store grader output
func setupLogFile(jobId string) (string, error, string) {
	outputFile := fmt.Sprintf("%s/%s.txt", OutputFolder.GetStr(), jobId)
	outFile, err := os.Create(outputFile)
	if err != nil {
		return "", err, fmt.Sprintf("error while creating log file at %s", outputFile)
	}
	defer func() {
		err := outFile.Close()
		if err != nil {
			log.Warn().Err(err).Msg("error while closing log file")
		}
	}()

	full, err := filepath.Abs(outputFile)
	if err != nil {
		return "", err, "error while getting absolute path"
	}

	return full, nil, ""
}

func (job *JobService) getLab(labId uint) (*models.Lab, error) {
	if labId == 0 {
		return nil, fmt.Errorf("invalid lab ID %d", labId)
	}
	labData, err := job.labSrv.GetLabFromDB(labId)
	if err != nil {
		return nil, err
	}
	return labData, nil
}

// zerolog log with context, intended for job logs
//
// shorthand for log.Ctx(ctx)
func jog(ctx context.Context) *zerolog.Logger {
	return log.Ctx(ctx)
}
