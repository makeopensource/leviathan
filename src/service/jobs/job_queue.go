package jobs

import (
	"context"
	"errors"
	"fmt"
	dk "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

type JobQueue struct {
	jobChannel chan *models.Job
	db         *gorm.DB
	dkSrv      *docker.DkService
	contextMap *models.Map[string, func()]
}

func NewJobQueue(totalJobs uint, db *gorm.DB, dk *docker.DkService) *JobQueue {
	queue := &JobQueue{
		jobChannel: make(chan *models.Job, totalJobs),
		contextMap: &models.Map[string, func()]{},
		db:         db,
		dkSrv:      dk,
	}

	queue.spawnWorkers(int(totalJobs))
	return queue
}

func (q *JobQueue) spawnWorkers(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go q.worker()
	}
}

func (q *JobQueue) AddJob(mes *models.Job) error {
	jog(mes.JobCtx).Info().Msg("sending job to queue")
	err := mes.ValidateForQueue()
	if err != nil {
		jog(mes.JobCtx).Err(err).Msg("job validation failed")
		return err
	}

	// run in go routine in case queue is full and gets blocked
	go func() {
		q.jobChannel <- mes
	}()
	return nil
}

func (q *JobQueue) NewJobContext(jobID string) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	wrapCancelFunc := func() {
		cancel()
		q.contextMap.Delete(jobID)
	}

	q.contextMap.Store(jobID, wrapCancelFunc)
	return common.CreateJobSubLoggerCtx(ctx, jobID)
}

func (q *JobQueue) CancelJob(messageId string) {
	cancel, ok := q.contextMap.Load(messageId)
	if !ok {
		log.Warn().Str(common.JobLogKey, messageId).Msg("job context was nil")
		return
	}
	cancel()
}

func (q *JobQueue) worker() {
	for msg := range q.jobChannel {
		if msg == nil {
			log.Error().Msg("job received was nil, THIS SHOULD NEVER HAPPEN")
			continue
		}

		if errors.Is(msg.JobCtx.Err(), context.Canceled) {
			q.setJobAsCancelled(msg)
			q.cleanupJob(msg, nil)
			jog(msg.JobCtx).Warn().Msgf("job context was canceled before queue could process")
			continue
		}

		jog(msg.JobCtx).Info().Msgf("worker is now processing job")
		q.runJob(msg)
	}
}

// runJob should ALWAYS BE BLOCKING, as it prevents the worker from moving on to a new job
func (q *JobQueue) runJob(job *models.Job) {
	client, contId, err, reason := q.setupJob(job)
	defer q.cleanupJob(job, client)
	if err != nil {
		q.bigProblem(job, reason, err)
		return
	}

	logStatusCh := make(chan struct {
		message string
		err     error
	})

	q.setJobInProgress(job)
	err = client.StartContainer(contId)
	if err != nil {
		q.bigProblem(job, "unable to start job container", err)
		return
	}

	// start writing to log file so that
	// we can stream changes to the log file to the user
	go func() {
		statusMessage, err2 := writeLogs(client, job)
		logStatusCh <- struct {
			message string
			err     error
		}{message: statusMessage, err: err2}
	}()

	statusCh, errCh := client.Client.ContainerWait(context.Background(), contId, dk.WaitConditionNotRunning)
	select {
	case <-statusCh:
		mes := <-logStatusCh
		if mes.err != nil {
			q.bigProblem(job, mes.message, mes.err)
			return
		}
		logLine, errMessage, err := verifyLogs(job)
		if err != nil || errMessage != "" {
			q.bigProblem(job, errMessage, err)
			return
		}
		q.greatSuccess(job, logLine)
		return
	case err := <-errCh:
		q.bigProblem(job, "error occurred while waiting for job process", err)
		return
	case <-time.After(job.JobTimeout):
		q.bigProblem(job, fmt.Sprintf("Maximum timeout reached for job, job ran for %s", job.JobTimeout), nil)
		return
	case <-job.JobCtx.Done():
		q.setJobAsCancelled(job)
		return
	}
}

// setupJob Set up job like king, yes!
// returns nil client if an error occurred while setup,
// make sure to handle null ptr dereference
func (q *JobQueue) setupJob(msg *models.Job) (*docker.DkClient, string, error, string) {
	q.setJobInSetup(msg)

	machine, err := q.dkSrv.ClientManager.GetClientById(msg.MachineId)
	if err != nil {
		return nil, "", err, "Failed to get machine info"
	}

	// incase dockerfile is not passed and referenced via tag name
	if msg.LabData.DockerFilePath != "" {
		err = machine.BuildImageFromDockerfile(msg.LabData.DockerFilePath, msg.LabData.ImageTag)
		if err != nil {
			return nil, "", err, "Failed to create image"
		}
		// folder structure is '/<id>/autolab/Dockerfile, get the <random id> folder path
		err = os.RemoveAll(filepath.Dir(filepath.Dir(msg.LabData.DockerFilePath)))
		if err != nil {
			return nil, "", err, "failed to delete dockerfile"
		}
	}

	resources := dk.Resources{
		NanoCPUs:  msg.JobLimits.NanoCPU * models.CPUQuota,
		Memory:    msg.JobLimits.Memory * models.MB,
		PidsLimit: &msg.JobLimits.PidsLimit,
	}

	contId, err := machine.CreateNewContainer(msg.JobId, msg.LabData.ImageTag, filepath.Base(msg.TmpJobFolderPath), msg.JobEntryCmd, resources)
	if err != nil {
		return nil, "", err, "unable to create job container"
	}

	err = machine.CopyToContainer(contId, msg.TmpJobFolderPath)
	if err != nil {
		return nil, "", err, "unable to copy files to job container"
	}

	msg.ContainerId = contId
	res := q.db.Save(msg)
	if res.Error != nil {
		return nil, "", err, "unable to update job in db"
	}

	return machine, contId, nil, ""
}

// cleanupJob clean up job,
// updates job in DB, removes the container and associated tmp job data
func (q *JobQueue) cleanupJob(msg *models.Job, client *docker.DkClient) {
	jog(msg.JobCtx).Info().Msg("cleaning up job")
	q.updateJobVeryNice(msg)

	if client != nil {
		err := client.RemoveContainer(msg.ContainerId, true, true)
		if err != nil {
			jog(msg.JobCtx).Warn().Err(err).Str("containerID", msg.ContainerId).Msg("unable to remove container")
		}
	}

	q.dkSrv.ClientManager.DecreaseJobCount(msg.MachineId)
	q.contextMap.Delete(msg.JobId)

	tmpFold := filepath.Dir(msg.TmpJobFolderPath) // get the dir above autolab subdir
	err := os.RemoveAll(tmpFold)
	if err != nil {
		jog(msg.JobCtx).Warn().Err(err).Str("dir", tmpFold).Msg("unable to remove tmp job directory")
		return
	}
}

// greatSuccess set job status to models.Complete
//
// Very nice!
//
// jobResult is the last line expected to be valid json string, returned to the job caller
func (q *JobQueue) greatSuccess(job *models.Job, jobResult string) {
	jog(job.JobCtx).Info().Msg("job completed successfully")
	job.Status = models.Complete
	job.StatusMessage = jobResult
}

// bigProblem set job status to models.Failed
//
// job failed, Not good!
//
// The publicReason will be displayed to the end user, providing a user-friendly message.
//
// The err parameter holds the underlying error, used for debugging purposes.
func (q *JobQueue) bigProblem(job *models.Job, publicReason string, err error) {
	jog(job.JobCtx).Error().Err(err).Str("reason", publicReason).Msg("job failed")
	job.Status = models.Failed
	job.StatusMessage = publicReason
	if err != nil {
		job.Error = err.Error()
	}
}

func (q *JobQueue) setJobAsCancelled(job *models.Job) {
	jog(job.JobCtx).Info().Msg("job was cancelled")
	job.Status = models.Canceled
	job.StatusMessage = "Job was cancelled"
}

// setJobInProgress set job status as models.Running
//
// Job is in progress, success soon!
func (q *JobQueue) setJobInProgress(msg *models.Job) {
	msg.Status = models.Running
	q.updateJobVeryNice(msg)
}

// setJobInSetup set job status as models.Preparing
//
// job is being setup standby
func (q *JobQueue) setJobInSetup(msg *models.Job) {
	msg.Status = models.Preparing
	q.updateJobVeryNice(msg)
}

// updateJobVeryNice Database updated, fresh like new wife!
func (q *JobQueue) updateJobVeryNice(msg *models.Job) {
	res := q.db.Save(msg)
	if res.Error != nil {
		jog(msg.JobCtx).Error().Err(res.Error).Msg("error occurred while saving job to db")
	}
}

func writeLogs(client *docker.DkClient, msg *models.Job) (string, error) {
	outputFile, err := os.OpenFile(msg.OutputLogFilePath, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return "unable to open log file", err
	}

	defer func() {
		err := outputFile.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error while closing output file")
		}
	}()

	logs, err := client.TailContainerLogs(context.Background(), msg.ContainerId)
	if err != nil {
		return "unable to tail job container", err
	}

	_, err = stdcopy.StdCopy(outputFile, outputFile, logs)
	if err != nil {
		return "unable to write to log file", err
	}
	return "", nil
}

func verifyLogs(msg *models.Job) (string, string, error) {
	if msg.Status == models.Failed {
		return "", "Job failed, skipping parsing log file", nil
	}

	outputFile, err := os.Open(msg.OutputLogFilePath)
	if err != nil {
		return "", "unable to open log file", err
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			jog(msg.JobCtx).Warn().Err(err).Msg("An error occurred while closing log file")
		}
	}(outputFile)

	line, err := common.GetLastLine(outputFile)
	if err != nil {
		return "", "unable to get logs", err
	}
	if !common.IsValidJSON(line) {
		return "", "unable to parse log output", nil
	}

	return line, "", nil
}
