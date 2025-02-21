package jobs

import (
	"context"
	"errors"
	"fmt"
	cond "github.com/docker/docker/api/types/container"
	cont "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type JobQueue struct {
	jobChannel chan *models.Job
	totalJobs  uint
	db         *gorm.DB
	dkSrv      *docker.DkService
	contextMap *models.Map[string, func()]
}

func NewJobQueue(totalJobs uint, db *gorm.DB, dk *docker.DkService) *JobQueue {
	queue := &JobQueue{
		jobChannel: make(chan *models.Job, totalJobs),
		totalJobs:  totalJobs,
		db:         db,
		dkSrv:      dk,
		contextMap: &models.Map[string, func()]{},
	}

	queue.CreateJobProcessors()
	return queue
}

func (q *JobQueue) CreateJobProcessors() {
	for i := 1; i < int(q.totalJobs); i++ {
		go q.worker(i)
	}
}

func (q *JobQueue) AddJob(mes *models.Job) error {
	err := mes.ValidateForQueue()
	if err != nil {
		jobLog(mes.JobCtx).Err(err).Msg("job validation failed")
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
	cancel := q.getJobCancelFunc(messageId)
	if cancel == nil {
		log.Warn().Str(common.JobLogKey, messageId).Msg("job context was nil")
		return
	}
	cancel()
}

func (q *JobQueue) getJobCancelFunc(messageId string) context.CancelFunc {
	val, ok := q.contextMap.Load(messageId)
	if !ok {
		return nil
	}
	return val
}

func (q *JobQueue) worker(workerId int) {
	for msg := range q.jobChannel {
		if errors.Is(msg.JobCtx.Err(), context.Canceled) {
			q.setJobAsCancelled(msg)
			jobLog(msg.JobCtx).Warn().Msgf("job context was canceled before queue could process")
			continue
		}

		jobLog(msg.JobCtx).Info().Int("workerID", workerId).Msgf("Worker is now processing job")
		q.runJob(msg)
	}
}

// runJob should ALWAYS BE BLOCKING, as it prevents the worker from moving on to a new job
func (q *JobQueue) runJob(msg *models.Job) {
	q.setJobInSetup(msg)

	client, contId, err, reason := q.setupJob(msg)
	defer q.cleanupJob(msg, client)

	if err != nil {
		q.bigProblem(reason, msg, err)
		return
	}

	q.setJobInProgress(msg)

	err = client.StartContainer(contId)
	if err != nil {
		q.bigProblem("unable to start job container", msg, err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	// start writing to log file so that
	// we can stream changes to the log file to the user
	go func() {
		defer wg.Done()
		q.writeLogs(client, msg)
	}()

	statusCh, errCh := client.Client.ContainerWait(context.Background(), contId, cond.WaitConditionNotRunning)
	select {
	case _ = <-statusCh:
		wg.Wait() // for logs to complete writing
		q.verifyLogs(msg.OutputLogFilePath, msg)
		return
	case err := <-errCh:
		q.bigProblem("error occurred while waiting for job process", msg, err)
		return
	case <-time.After(msg.JobTimeout):
		q.bigProblem(fmt.Sprintf("Maximum timeout reached for job, job ran for %s", msg.JobTimeout), msg, nil)
		return
	case <-msg.JobCtx.Done():
		q.setJobAsCancelled(msg)
		return
	}
}

func (q *JobQueue) writeLogs(client *docker.DkClient, msg *models.Job) {
	outPutFile, err := os.OpenFile(msg.OutputLogFilePath, os.O_RDWR|os.O_CREATE, 660)
	if err != nil {
		q.bigProblem("unable to open output file", msg, err)
		return
	}

	defer func() {
		err := outPutFile.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error while closing output file")
		}
	}()

	logs, err := client.TailContainerLogs(context.Background(), msg.ContainerId)
	if err != nil {
		q.bigProblem("unable to tail job container", msg, err)
		return
	}

	_, err = stdcopy.StdCopy(outPutFile, outPutFile, logs)
	if err != nil {
		q.bigProblem("unable to write to output file", msg, err)
		return
	}
}

// setupJob Set up job like king, yes!
// returns nil client if an error occurred while setup,
// make sure to handle null ptr dereference
func (q *JobQueue) setupJob(msg *models.Job) (*docker.DkClient, string, error, string) {
	machine, err := q.dkSrv.ClientManager.GetClientById(msg.MachineId)
	if err != nil {
		return nil, "", nil, "Failed to get machine info"
	}

	// incase dockerfile is not passed and referenced via tag name
	if msg.LabData.DockerFilePath != "" {
		err = machine.BuildImageFromDockerfile(msg.LabData.DockerFilePath, msg.LabData.ImageTag)
		if err != nil {
			return nil, "", err, "Failed to create image"
		}
		err = os.Remove(msg.LabData.DockerFilePath)
		if err != nil {
			return nil, "", err, "failed to delete dockerfile"
		}
	}

	pidsLimit := int64(100)
	// todo load from job message
	resources := cont.Resources{
		NanoCPUs:  1 * models.CPUQuota, // 1 virtual core
		Memory:    512 * models.MB,
		PidsLimit: &pidsLimit,
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

// bigProblem job failed, Not good!
// The publicReason will be displayed to the end user, providing a user-friendly message.
// The err parameter holds the underlying error, used for debugging purposes.
// we use variadic param to make jobtar optional
func (q *JobQueue) bigProblem(publicReason string, job *models.Job, err error) {
	jobLog(job.JobCtx).Error().Err(err).Str("reason", publicReason).Msg("job failed")

	// job will be saved by the cleanup function
	job.Status = models.Failed
	job.StatusMessage = publicReason
}

func (q *JobQueue) setJobAsCancelled(job *models.Job) {
	jobLog(job.JobCtx).Info().Msg("job was cancelled")
	job.Status = models.Canceled
	job.StatusMessage = "Job was cancelled"
}

// greatSuccess Very nice!
// jobResult is the last line expected to be valid json string, returned to the job caller
func (q *JobQueue) greatSuccess(job *models.Job, jobResult string) {
	jobLog(job.JobCtx).Info().Msg("job completed successfully")
	job.Status = models.Complete
	job.StatusMessage = jobResult
}

// cleanupJob clean up job
// sets job to success, removes the container and associated tmp job data
func (q *JobQueue) cleanupJob(msg *models.Job, client *docker.DkClient) {
	jobLog(msg.JobCtx).Info().Msg("cleaning up job")

	q.updateJobVeryNice(msg)

	if client != nil {
		err := client.RemoveContainer(msg.ContainerId, true, true)
		if err != nil {
			jobLog(msg.JobCtx).Warn().Err(err).Str("containerID", msg.ContainerId).Msg("unable to remove container")
		}
	}

	q.dkSrv.ClientManager.DecreaseJobCount(msg.MachineId)

	tmpFold := filepath.Dir(msg.TmpJobFolderPath) // get the dir above autolab subdir
	err := os.RemoveAll(tmpFold)
	if err != nil {
		jobLog(msg.JobCtx).Warn().Err(err).Str("dir", tmpFold).Msg("unable to remove tmp job directory")
		return
	}
}

// Job is in progress, success soon!
func (q *JobQueue) setJobInProgress(msg *models.Job) {
	msg.Status = models.Running
	q.updateJobVeryNice(msg)
}

// job is being setup standby
func (q *JobQueue) setJobInSetup(msg *models.Job) {
	msg.Status = models.Preparing
	q.updateJobVeryNice(msg)
}

// updateJobVeryNice Database updated, fresh like new wife!
func (q *JobQueue) updateJobVeryNice(msg *models.Job) {
	res := q.db.Save(msg)
	if res.Error != nil {
		jobLog(msg.JobCtx).Error().Err(res.Error).Msg("error occurred while saving job to db")
	}
}

func (q *JobQueue) verifyLogs(file string, msg *models.Job) {
	if msg.Status == models.Failed {
		jobLog(msg.JobCtx).Warn().Msg("Job failed, skipping parsing log file")
		return
	}

	outputFile, err := os.Open(file)
	if err != nil {
		q.bigProblem("unable to open log file", msg, err)
		return
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			jobLog(msg.JobCtx).Warn().Err(err).Msg("An error occurred while closing log file")
		}
	}(outputFile)

	line, err := common.GetLastLine(outputFile)
	if err != nil {
		q.bigProblem("unable to get logs", msg, err)
		return
	}

	if !common.IsValidJSON(line) {
		q.bigProblem("unable to parse log output", msg, err)
		return
	}

	q.greatSuccess(msg, line)
}
