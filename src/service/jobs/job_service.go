package jobs

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/google/uuid"
	com "github.com/makeopensource/leviathan/common"
	v2 "github.com/makeopensource/leviathan/generated/jobs/v1"
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

// WaitForJobAndLogs blocks until job reaches, Cancelled, Complete, error
func (job *JobService) WaitForJobAndLogs(ctx context.Context, jobUuid string) (*models.Job, string, error) {
	jobInf, complete, content, err := job.checkJob(jobUuid)
	if err != nil {
		return nil, "", err
	}
	if complete {
		return jobInf, content, nil
	}

	jobInfoCh := job.SubToJob(jobUuid)
	defer job.UnsubToJob(jobUuid)

	logContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	logsCh := job.ListenToJobLogs(logContext, jobInf)
	if err != nil {
		return nil, "", err
	}

	var jobInfo *models.Job
	var logs string
	var jobOk = false

	// Keep listening until channel closes, implying job is complete
	for {
		if jobOk {
			// read log file one last time
			content := com.ReadLogFile(jobInfo.OutputLogFilePath)
			return jobInfo, content, nil
		}

		select {
		case logsTmp, ok := <-logsCh:
			if ok && logs != logsTmp {
				logs = logsTmp
			}
		case jobTmp, ok := <-jobInfoCh:
			jobOk = !ok
			if ok {
				jobInfo = jobTmp
			}
		case <-ctx.Done():
			return jobInfo, logs, fmt.Errorf("context canceled")
		}
	}
}

func (job *JobService) StreamJobAndLogs(ctx context.Context, jobUuid string, stream *connect.ServerStream[v2.JobLogsResponse]) error {
	jobInfo, complete, flogs, err := job.checkJob(jobUuid)
	if err != nil {
		return err
	}

	// send initial job data
	if jobInfo != nil {
		err := sendJobToStream(stream, jobInfo, flogs)
		if err != nil {
			return err
		}
	}
	if complete {
		return nil
	}

	jobInfoCh := job.SubToJob(jobInfo.JobId)
	defer job.UnsubToJob(jobUuid)

	// since ListenToJobLogs reads logs in an infinite loop,
	// always cancel this context, exiting the loop
	logContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	logsCh := job.ListenToJobLogs(logContext, jobInfo)
	if err != nil {
		return err
	}

	var logs string
	var jobOk = false

	// Keep listening until channel closes, indicating job is complete
	for {
		if jobOk {
			content := com.ReadLogFile(jobInfo.OutputLogFilePath)
			if err != nil {
				log.Warn().Err(err).Str("path", jobInfo.OutputLogFilePath).Msg("Failed to read job log file")
			}
			err = sendJobToStream(stream, jobInfo, content)
			if err != nil {
				return err
			}
			// job done
			return nil
		}

		select {
		case logsTmp, ok := <-logsCh:
			if ok {
				log.Debug().Msg("job logs changed")
				logs = logsTmp
				err := sendJobToStream(stream, jobInfo, logsTmp)
				if err != nil {
					return err
				}
			}
		case jobTmp, ok := <-jobInfoCh:
			jobOk = !ok
			if ok && jobInfo != nil && jobTmp.Status != jobInfo.Status {
				log.Debug().Msgf("job status changed from %s to %s", jobInfo.Status, jobTmp.Status)
				jobInfo = jobTmp
				err := sendJobToStream(stream, jobInfo, logs)
				if err != nil {
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
	go func() {
		prevLength := 0
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		// keep reading until ctx is done
		for {
			select {
			case <-ticker.C:
				content := com.ReadLogFile(jobInfo.OutputLogFilePath)
				// send if content changed
				if len(content) > prevLength {
					log.Debug().Str(com.JobLogKey, jobInfo.JobId).Msgf("sending log, length changed from %d to %d", prevLength, len(content))
					prevLength = len(content)
					logChannel <- content
				}
			case <-ctx.Done():
				log.Debug().Str(com.JobLogKey, jobInfo.JobId).Msg("stopping listening for logs")
				return
			}
		}
	}()

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

func sendJobToStream(stream *connect.ServerStream[v2.JobLogsResponse], jobInfo *models.Job, logs string) error {
	err := stream.Send(&v2.JobLogsResponse{
		JobInfo: &v2.JobStatus{
			JobId:         jobInfo.JobId,
			Status:        string(jobInfo.Status),
			StatusMessage: jobInfo.StatusMessage,
		},
		Logs: logs,
	})
	if err != nil {
		return err
	}
	return nil
}

// zerolog log with context, intended for job logs
//
// shorthand for log.Ctx(ctx)
func jog(ctx context.Context) *zerolog.Logger {
	return log.Ctx(ctx)
}
