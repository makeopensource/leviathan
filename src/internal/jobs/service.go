package jobs

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/internal/config"
	"github.com/makeopensource/leviathan/internal/docker"
	fm "github.com/makeopensource/leviathan/internal/file_manager"
	"github.com/makeopensource/leviathan/internal/labs"
	fu "github.com/makeopensource/leviathan/pkg/file_utils"
	"github.com/makeopensource/leviathan/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"time"
)

type JobService struct {
	dockerSrv   *docker.DkService
	queue       *JobQueue
	broadcastCh *BroadcastChannel
	labSrv      *labs.LabService
	fileManSrv  *fm.FileManagerService
	db          JobStore
}

func NewJobService(
	db JobStore,
	bc *BroadcastChannel,
	dockerService *docker.DkService,
	labService *labs.LabService,
	tmpFileService *fm.FileManagerService,
) *JobService {
	queue := NewJobQueue(
		uint(config.ConcurrentJobs.GetUint64()),
		db,
		dockerService,
		labService,
	)

	srv := &JobService{
		db:          db,
		broadcastCh: bc,
		dockerSrv:   dockerService,
		queue:       queue,
		labSrv:      labService,
		fileManSrv:  tmpFileService,
	}
	srv.cleanupOrphanJobs()
	return srv
}

func (job *JobService) NewJob(newJob *Job, submissionFolderId string) (string, error) {
	if newJob.LabData == nil {
		labData, err := job.getLab(newJob.LabID)
		if err != nil {
			return "", err
		}
		newJob.LabData = labData
	}

	jobId, err := uuid.NewUUID()
	if err != nil {
		return "", logger.ErrLog("failed to generate job ID", err, log.Error())
	}
	newJob.JobId = jobId.String()

	mId := job.dockerSrv.ClientManager.GetLeastJobCountMachineId()

	// job context, so that it can be cancelled, and store sub logger
	ctx := job.queue.NewJobContext(newJob.JobId)

	jobDir, err := CreateTmpJobDir(newJob.JobId, config.SubmissionFolder.GetStr())
	if err != nil {
		return "", logger.ErrLog("failed to create job dir", err, jog(ctx).Error())
	}

	submissionFolder, err := job.fileManSrv.GetSubmissionPath(submissionFolderId)
	if err != nil {
		return "", err
	}
	defer job.fileManSrv.DeleteFolder(submissionFolderId)

	if err = fu.HardLinkFolder(submissionFolder, jobDir); err != nil {
		return "", logger.ErrLog("unable to copy files to job dir", err, log.Error())
	}
	if err = fu.HardLinkFolder(newJob.LabData.JobFilesDirPath, jobDir); err != nil {
		return "", logger.ErrLog("unable to copy files to job dir", err, log.Error())
	}

	logPath, err2 := setupLogFile(newJob.JobId)
	if err2 != nil {
		return "", logger.ErrLog(
			"failed to setup log file: "+err2.Reason(),
			err,
			jog(ctx).Error().Str("reason", err2.Reason()),
		)
	}

	// setup job metadata
	newJob.MachineId = mId
	newJob.Status = Queued
	newJob.OutputLogFilePath = logPath
	newJob.TmpJobFolderPath = jobDir
	newJob.JobCtx = ctx
	newJob.LabData.VerifyJobLimits()
	jog(newJob.JobCtx).Debug().Any("limits", newJob.LabData.JobLimits).Msg("job limits")

	err = job.db.CreateJob(newJob)
	if err != nil {
		return "", logger.ErrLog("failed to save job to db", err, jog(ctx).Error())
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
func (job *JobService) WaitForJobAndLogs(jobUuid string) (*Job, string, error) {
	var jobInfo *Job
	var outerLogs string

	err := job.StreamJobAndLogs(
		context.Background(),
		jobUuid,
		func(jobInf *Job, logs string) error {
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
	streamFunc func(jobInf *Job, logs string) error,
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
func (job *JobService) GetJobStatusAndLogs(jobUuid string) (*Job, string, error) {
	jobInfo, _, logs, err := job.checkJob(jobUuid)
	if err != nil {
		return nil, "", err
	}
	return jobInfo, logs, nil
}

func (job *JobService) ListenToJobLogs(ctx context.Context, jobInfo *Job) chan string {
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
					log.Debug().Str(logger.JobLogKey, jobInfo.JobId).Msgf("sending log, length changed from %d to %d", prevLength, contLen)
					prevLength = contLen
					logChannel <- content
				}
			case <-ctx.Done():
				close(logChannel)
				log.Debug().Str(logger.JobLogKey, jobInfo.JobId).Msg("stopping listening for logs")
				return
			}
		}
	}(ctx)

	return logChannel
}

func (job *JobService) SubToJob(jobUuid string) chan *Job {
	return job.broadcastCh.Subscribe(jobUuid)
}

func (job *JobService) UnsubToJob(jobUuid string) {
	job.broadcastCh.Unsubscribe(jobUuid)
}

func (job *JobService) checkJob(jobUuid string) (*Job, bool, string, error) {
	jobInf, err := job.GetJobFromDB(jobUuid)
	if err != nil {
		return nil, false, "", fmt.Errorf("failed to get job info")
	}

	if jobInf.Status.Done() {
		log.Debug().Str(logger.JobLogKey, jobUuid).Msg("job is already done")

		content := ReadLogFile(jobInf.OutputLogFilePath)
		return jobInf, true, content, nil
	} else {
		return jobInf, false, "", nil
	}
}

func (job *JobService) GetJobFromDB(jobUuid string) (*Job, error) {
	jobInfo, err := job.db.GetJobByUuid(jobUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get job info from db: %v", err)
	}
	return jobInfo, nil
}

// removes any job left in an 'active' state before application start,
// fail any jobs that were running before leviathan was able to process them (for whatever reason)
//
// for example machine running leviathan shutdown unexpectedly or leviathan had an unrecoverable error
func (job *JobService) cleanupOrphanJobs() {
	var orphanJobs []Job
	err := job.db.FetchInProgressJobs(&orphanJobs)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to query database for orphan jobs")
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

		orphan.Status = Failed
		orphan.StatusMessage = "job was unable to be processed due to an internal server error"
		err = job.db.UpdateJob(&orphan)
		if err != nil {
			log.Warn().Err(err).Msg("unable to update orphan job status")
		}
	}

	if len(orphanJobs) != 0 {
		log.Info().Msgf("Cleaned up %d orphan jobs", len(orphanJobs))
	}
}

// setupLogFile store grader output
func setupLogFile(jobId string) (string, JobError) {
	outputFile := fmt.Sprintf("%s/%s.txt", config.OutputFolder.GetStr(), jobId)
	outFile, err := os.Create(outputFile)
	if err != nil {
		return "", JError(fmt.Sprintf("error while creating log file at %s", outputFile), err)
	}
	defer func() {
		err := outFile.Close()
		if err != nil {
			log.Warn().Err(err).Msg("error while closing log file")
		}
	}()

	full, err := filepath.Abs(outputFile)
	if err != nil {
		return "", JError("error while getting absolute path", err)
	}

	return full, nil
}

func (job *JobService) getLab(labId uint) (*labs.Lab, error) {
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
