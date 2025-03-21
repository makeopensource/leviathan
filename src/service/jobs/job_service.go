package jobs

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	com "github.com/makeopensource/leviathan/common"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
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
}

func NewJobService(db *gorm.DB, bc *models.BroadcastChannel, dockerService *docker.DkService) *JobService {
	return &JobService{
		db:          db,
		broadcastCh: bc,
		dockerSrv:   dockerService,
		queue:       NewJobQueue(uint(com.ConcurrentJobs.GetUint64()), db, dockerService),
	}
}

func (job *JobService) NewJob(newJob *models.Job, jobFiles []*v1.FileUpload, dockerfile *v1.FileUpload) (string, error) {
	jobId, err := uuid.NewUUID()
	if err != nil {
		log.Error().Err(err).Msg("failed to generate job ID")
		return "", fmt.Errorf("failed to generate job ID")
	}
	newJob.JobId = jobId.String()

	mId := job.dockerSrv.ClientManager.GetLeastJobCountMachineId()

	// job context, so that it can be cancelled, and store sub logger
	ctx := job.queue.NewJobContext(newJob.JobId)

	// "" implies randomly named tmp folder
	dockerFileDir, err := com.CreateTmpJobDir(newJob.JobId, "", dockerfile)
	if err != nil {
		jog(ctx).Error().Err(err).Msg("failed to create dockerfile dir")
		return "", fmt.Errorf("failed to create dockerfile dir")
	}

	jobDir, err := com.CreateTmpJobDir(newJob.JobId, com.SubmissionFolder.GetStr(), jobFiles...)
	if err != nil {
		jog(ctx).Error().Err(err).Msg("failed to create job dir")
		return "", fmt.Errorf("failed to create job dir")
	}

	logPath, err, reason := setupLogFile(newJob.JobId)
	if err != nil {
		jog(ctx).Error().Err(err).Str("reason", reason).Msg("failed to setup log file")
		return "", fmt.Errorf("failed to setup log file")
	}

	// setup job metadata
	newJob.MachineId = mId
	newJob.Status = models.Queued
	newJob.OutputLogFilePath = logPath
	newJob.TmpJobFolderPath = jobDir
	newJob.LabData.DockerFilePath = fmt.Sprintf("%s/%s", dockerFileDir, dockerfile.Filename)
	newJob.JobCtx = ctx
	newJob.VerifyJobLimits()
	jog(newJob.JobCtx).Debug().Any("limits", newJob.JobLimits).Msg("job limits")

	res := job.db.Create(newJob)
	if res.Error != nil {
		jog(ctx).Error().Err(res.Error).Msg("Failed to save job to db")
		return "", fmt.Errorf("failed to save job to db")
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
			cancel()                                              // stop updating log channel
			content := com.ReadLogFile(jobInfo.OutputLogFilePath) // final read from the log file
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
				content := com.ReadLogFile(jobInfo.OutputLogFilePath)
				contLen := len(content)
				if contLen > prevLength { // send if content changed
					log.Debug().Str(com.JobLogKey, jobInfo.JobId).Msgf("sending log, length changed from %d to %d", prevLength, contLen)
					prevLength = contLen
					logChannel <- content
				}
			case <-ctx.Done():
				close(logChannel)
				log.Debug().Str(com.JobLogKey, jobInfo.JobId).Msg("stopping listening for logs")
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
		log.Debug().Str(com.JobLogKey, jobUuid).Msg("job is already done")

		content := com.ReadLogFile(jobInf.OutputLogFilePath)
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

// setupLogFile store grader output
func setupLogFile(jobId string) (string, error, string) {
	outputFile := fmt.Sprintf("%s/%s.txt", com.OutputFolder.GetStr(), jobId)
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

// zerolog log with context, intended for job logs
//
// shorthand for log.Ctx(ctx)
func jog(ctx context.Context) *zerolog.Logger {
	return log.Ctx(ctx)
}
