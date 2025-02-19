package jobs

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/common"
	v2 "github.com/makeopensource/leviathan/generated/jobs/v1"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

type JobService struct {
	db           *gorm.DB
	labFileCache *models.LabFilesCache
	dockerSrv    *docker.DkService
	queue        *JobQueue
	broadcastCh  *models.BroadcastChannel
}

func NewJobService(db *gorm.DB, cache *models.LabFilesCache, bc *models.BroadcastChannel, dockerService *docker.DkService) *JobService {
	return &JobService{
		db:           db,
		broadcastCh:  bc,
		labFileCache: cache,
		dockerSrv:    dockerService,
		queue:        NewJobQueue(uint(common.ConcurrentJobs.GetUint64()), db, dockerService),
	}
}

func (job *JobService) NewJob(newJob *models.Job, makefile *v1.FileUpload, grader *v1.FileUpload, student *v1.FileUpload, dockerfile *v1.FileUpload) (string, error) {
	mId := job.dockerSrv.ClientManager.GetLeastJobCountMachineId()
	if mId == "" {
		return "", fmt.Errorf("failed to assign machine")
	}

	jobId, err := uuid.NewUUID()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to generate job ID")
		return "", fmt.Errorf("failed to generate job ID")
	}

	jobDir, err := common.CreateTmpJobDir(
		jobId.String(),
		map[string][]byte{
			grader.Filename:     grader.Content,
			makefile.Filename:   makefile.Content,
			student.Filename:    student.Content,
			dockerfile.Filename: dockerfile.Content,
		},
	)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to create job dir")
		return "", fmt.Errorf("failed to create job dir")
	}

	// setup job metadata
	newJob.JobId = jobId.String()
	newJob.MachineId = mId
	newJob.Status = models.Queued
	newJob.OutputLogFilePath = job.setupLogFile(newJob.JobId)
	newJob.TmpJobFolderPath = jobDir
	newJob.LabData.DockerFilePath = fmt.Sprintf("%s/%s", newJob.TmpJobFolderPath, dockerfile.Filename)
	// job context, so that it can be cancelled
	newJob.JobCtx = job.queue.NewJobContext(newJob.JobId)

	res := job.db.Create(newJob)
	if res.Error != nil {
		log.Error().Err(res.Error).Msgf("Failed to save job to db")
		return "", fmt.Errorf("failed to save job to db")
	}

	// run in go routine in case queue is full and this gets blocked
	go job.queue.AddJob(newJob)

	return jobId.String(), nil
}

func (job *JobService) CancelJob(jobUuid string) {
	job.queue.CancelJob(jobUuid)
}

// WaitForJob is similar to GetJobStatus but is blocking and runs until job is complete or errors
func (job *JobService) WaitForJob(ctx context.Context, jobUuid string) (*models.Job, error) {
	jobInfoCh := job.SubToJob(jobUuid)
	defer func() {
		go job.UnsubToJob(jobUuid)
	}()

	var jobInfo *models.Job
	for {
		select {
		case tmp, ok := <-jobInfoCh:
			//log.Debug().Any("job", tmp).Msg("New job data")
			if !ok {
				return jobInfo, nil
			} else {
				jobInfo = tmp
			}
		case <-ctx.Done():
			log.Debug().Msg("Context done")
			return nil, fmt.Errorf("context canceled")
		}
	}
}

// WaitForJobAndLogs this is an experimental function, it is not working
// todo logs return empty
func (job *JobService) WaitForJobAndLogs(ctx context.Context, jobUuid string) (*models.Job, string, error) {
	jobInf, complete, content, err := job.checkJob(jobUuid)
	if err != nil {
		return nil, "", err
	}
	if complete {
		// job was returned, indicating job is complete
		return jobInf, content, nil
	}

	jobInfoCh := job.SubToJob(jobUuid)
	defer func() {
		go job.UnsubToJob(jobUuid)
	}()

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
			content := readLogFile(jobInfo)
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
	defer func() {
		go job.UnsubToJob(jobUuid)
	}()

	logContext, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		log.Debug().Msg("log context done")
	}()
	logsCh := job.ListenToJobLogs(logContext, jobInfo)
	if err != nil {
		return err
	}

	var logs string
	var jobOk = false

	// Keep listening until channel closes, indicating job is complete
	for {
		if jobOk {
			content := readLogFile(jobInfo)
			if err != nil {
				log.Warn().Err(err).Msgf("Failed to read job log file at %s", jobInfo.OutputLogFilePath)
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
			// stream on status change
			if ok && logsTmp != logs {
				log.Debug().Msgf("job logs changed")
				logs = logsTmp
				err := sendJobToStream(stream, jobInfo, logs)
				if err != nil {
					return err
				}
			}
		case jobTmp, ok := <-jobInfoCh:
			jobOk = !ok
			// stream on status change
			if ok && jobInfo != nil && jobTmp.Status != jobInfo.Status {
				log.Debug().Msgf("job Info changed")
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
	logChannel := make(chan string, 50)
	go func() {
		// keep reading until ctx is done
		for {
			// Read all the content of the file
			content := readLogFile(jobInfo)
			logChannel <- content

			// err if context was cancelled, i.e. connection closed
			if ctx.Err() != nil {
				log.Debug().Msgf("Stopping listening for logs: %s", jobInfo.JobId)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}()

	return logChannel
}

func (job *JobService) SubToJob(jobUuid string) chan *models.Job {
	jch := job.broadcastCh.Subscribe(jobUuid)
	return jch
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
		log.Debug().Msgf("Job %s is done", jobUuid)

		content, err := os.ReadFile(jobInf.OutputLogFilePath)
		if err != nil {
			log.Warn().Err(err).Msgf("Failed to read job log file at %s", jobInf.OutputLogFilePath)
			return nil, false, "", fmt.Errorf("failed to read log file")
		}

		return jobInf, true, string(content), nil
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
// this is blocking operation make sure to
// stream logs in a go routine
func (job *JobService) setupLogFile(jobId string) string {
	outputFile := fmt.Sprintf("%s/%s.txt", common.OutputFolder.GetStr(), jobId)
	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Error().Err(err).Msgf("Error while creating file")
		return ""
	}
	defer func() {
		err := outFile.Close()
		if err != nil {
			log.Error().Err(err).Msgf("Error while closing file")
		}
	}()

	full, err := filepath.Abs(outputFile)
	if err != nil {
		log.Error().Err(err).Msgf("Error while getting absolute path")
		return ""
	}

	return full
}

func sendJobToStream(stream *connect.ServerStream[v2.JobLogsResponse], jobInfo *models.Job, logs string) error {
	err := stream.Send(&v2.JobLogsResponse{
		JobInfo: &v2.JobStatus{
			JobId:            jobInfo.JobId,
			MachineId:        jobInfo.MachineId,
			ContainerId:      jobInfo.ContainerId,
			Status:           string(jobInfo.Status),
			StatusMessage:    jobInfo.StatusMessage,
			OutputFilePath:   "",
			TmpJobFolderPath: "",
			JobTimeout:       int64(jobInfo.JobTimeout),
		},
		Logs: logs,
	})
	if err != nil {
		return err
	}
	return nil
}

func readLogFile(jobInfo *models.Job) string {
	content, err := os.ReadFile(jobInfo.OutputLogFilePath)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to read job log file at %s", jobInfo.OutputLogFilePath)
		return ""
	}
	return string(content)
}
