package jobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
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

func (job *JobService) NewJob(jobReq *models.Job) (string, error) {
	jobId, err := uuid.NewUUID()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to generate job ID")
		return "", fmt.Errorf("failed to generate job ID")
	}

	tmp := job.dockerSrv.ClientManager.GetLeastJobCountMachineId()
	if tmp == "" {
		return "", fmt.Errorf("failed to assign machine")
	}

	jobReq.JobId = jobId.String()
	jobReq.MachineId = tmp
	jobReq.Status = models.Queued

	job.setupLogFile(jobReq)

	res := job.db.Create(jobReq)
	if res.Error != nil {
		log.Error().Err(res.Error).Msgf("Failed to save job to db")
		return "", fmt.Errorf("failed to save job to db")
	}

	// job context, so that it can be cancelled
	jobReq.JobCtx = job.queue.NewJobContext(jobReq.JobId)
	// run in go routine in case queue is full and this gets blocked
	go job.queue.AddJob(jobReq)

	return jobId.String(), nil
}

// WaitForJob is similar to GetJobStatus but is blocking and runs until job is complete or errors
func (job *JobService) WaitForJob(ctx context.Context, jobUuid string) (*models.Job, error) {
	jobInfoCh, _ := job.SubToJob(jobUuid)
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
	jobInfoCh, err := job.SubToJob(jobUuid)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		go job.UnsubToJob(jobUuid)
	}()

	var jobInfo *models.Job
	var logs string

	// First phase: Wait for job to start
	for {
		select {
		case jobTmp, ok := <-jobInfoCh:
			if !ok {
				return jobInfo, logs, nil
			} else {
				jobInfo = jobTmp
			}
			// Once job is running, break out and start log monitoring
			if jobInfo.Status == models.Running {
				goto monitorWithLogs
			}
		case <-ctx.Done():
			return jobInfo, logs, nil
		}
	}

monitorWithLogs:
	// Start log monitoring once job is active
	logsCh, errCh, err := job.ListenToJobLogs(ctx, jobUuid)
	if err != nil {
		return nil, "", err
	}

	var jobOk = false
	var logOk = false

	// Keep listening until complete
	for {
		if jobOk && logOk {
			return jobInfo, logs, nil
		}

		select {
		case logsTmp, ok := <-logsCh:
			logOk = !ok
			if ok {
				logs += logsTmp
			}
		case jobTmp, ok := <-jobInfoCh:
			jobOk = !ok
			if ok {
				jobInfo = jobTmp
			}
		case err = <-errCh:
			return jobInfo, logs, err
		case <-ctx.Done():
			return jobInfo, logs, fmt.Errorf("context canceled")
		}
	}
}

func (job *JobService) ListenToJobLogs(ctx context.Context, jobUuid string) (chan string, chan error, error) {
	jobInfo, err := job.getJob(jobUuid)
	if err != nil {
		return nil, nil, err
	}

	machine, err := job.dockerSrv.ClientManager.GetClientById(jobInfo.MachineId)
	if err != nil {
		return nil, nil, err
	}

	logs, err := machine.TailContainerLogs(ctx, jobInfo.ContainerId)
	if err != nil {
		return nil, nil, err
	}

	logChannel := make(chan string)
	errChannel := make(chan error)

	streamWriter := &models.LogChannelWriter{Channel: logChannel}
	go func() {
		_, err = stdcopy.StdCopy(streamWriter, streamWriter, logs)
		if err != nil && err != io.EOF && !errors.Is(err, context.Canceled) {
			log.Error().Err(err).Msgf("failed to tail logs for container")
			errChannel <- err
		}

		// done reading logs
		close(logChannel)
	}()

	return logChannel, errChannel, err
}

func (job *JobService) SubToJob(jobUuid string) (chan *models.Job, error) {
	jch := job.broadcastCh.Subscribe(jobUuid)

	// check job data
	getJob, err := job.getJob(jobUuid)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to get job")
		return nil, fmt.Errorf("unable to find job")
	}

	// send initial job data
	jch <- getJob
	return jch, nil
}

func (job *JobService) UnsubToJob(jobUuid string) {
	job.broadcastCh.Unsubscribe(jobUuid)
}

func (job *JobService) getJob(jobUuid string) (*models.Job, error) {
	var jobInfo *models.Job
	res := job.db.First(&jobInfo, "job_id = ?", jobUuid)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get job info from db")
	}
	return jobInfo, nil
}

func (job *JobService) CancelJob(jobUuid string) {
	job.queue.CancelJob(jobUuid)
}

// setupLogFile store grader output
// this is blocking operation make sure to
// stream logs in a go routine
func (job *JobService) setupLogFile(msg *models.Job) *os.File {
	outputFile := fmt.Sprintf("%s/%s.txt", common.OutputFolder.GetStr(), msg.JobId)
	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Error().Err(err).Msgf("Error while creating file")
		return nil
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
		return nil
	}

	msg.OutputFilePath = full
	return outFile
}
