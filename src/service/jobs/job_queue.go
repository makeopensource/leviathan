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

func (q *JobQueue) AddJob(mes *models.Job) {
	q.jobChannel <- mes
}

func (q *JobQueue) NewJobContext(messageId string) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	wrapCancelFunc := func() {
		cancel()
		q.contextMap.Delete(messageId)
	}

	q.contextMap.Store(messageId, wrapCancelFunc)
	return ctx
}

func (q *JobQueue) GetJobCancelFunc(messageId string) context.CancelFunc {
	val, ok := q.contextMap.Load(messageId)
	if !ok {
		return nil
	}
	return val
}

func (q *JobQueue) CancelJob(messageId string) {
	cancel := q.GetJobCancelFunc(messageId)
	if cancel == nil {
		log.Warn().Str("messageId", messageId).Msg("nil job context")
		return
	}
	cancel()
}

func (q *JobQueue) worker(workerId int) {
	for msg := range q.jobChannel {
		if errors.Is(msg.JobCtx.Err(), context.Canceled) {
			log.Warn().Msgf("job: %s context was canceled", msg.JobId)
			q.setJobAsCancelled(msg)
			continue
		}

		log.Info().Msgf("Worker: %d is now processing job: %d", workerId, msg.ID)
		q.runJob(msg)
	}
}

// runJob should ALWAYS BE BLOCKING, as it prevents the worker from moving on to a new job
func (q *JobQueue) runJob(msg *models.Job) {
	q.setJobInSetup(msg)

	client, contId, jobFolderPath, err := q.setupJob(msg)
	if err != nil {
		// error handled by setupJob
		return
	}
	defer q.cleanupJob(msg, client, jobFolderPath)

	q.setJobInProgress(msg)

	err = client.StartContainer(contId)
	if err != nil {
		q.bigProblem("Unable to start job container", msg, err)
		return
	}

	statusCh, errCh := client.Client.ContainerWait(context.Background(), contId, cond.WaitConditionNotRunning)
	select {
	case state := <-statusCh:
		log.Info().Any("state", state).Msgf("Job: %s completed", msg.JobId)
		q.writeLogs(client, msg)
		q.verifyLogs(msg.OutputFilePath, msg)
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
	outPutFile, err := os.OpenFile(msg.OutputFilePath, os.O_RDWR|os.O_CREATE, 660)
	if err != nil {
		q.bigProblem("Unable to open output file", msg, err)
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
		q.bigProblem("Unable to tail job container", msg, err)
		return
	}

	_, err = stdcopy.StdCopy(outPutFile, outPutFile, logs)
	if err != nil {
		q.bigProblem("unable to write to output file", msg, err)
		return
	}
}

// setupJob Set up job like king, yes!
// returns nil client if an error occurred while setup
func (q *JobQueue) setupJob(msg *models.Job) (*docker.DkClient, string, string, error) {
	jobDir, err := common.CreateTmpJobDir(
		msg.JobId,
		map[string][]byte{
			msg.LabData.GraderFilename:    msg.LabData.GraderFile,
			msg.LabData.MakeFilename:      msg.LabData.MakeFile,
			msg.StudentSubmissionFileName: msg.StudentSubmissionFile,
		},
	)
	if err != nil {
		q.bigProblem("Failed to convert job data to tar", msg, err, jobDir)
		return nil, "", "", nil
	}

	machine, err := q.dkSrv.ClientManager.GetClientById(msg.MachineId)
	if err != nil {
		q.bigProblem("Failed to get machine info", msg, err, jobDir)
		return nil, "", "", nil
	}

	pidsLimit := int64(100)
	// todo load from job message
	resources := cont.Resources{
		NanoCPUs:  1 * models.CPUQuota, // 1 virtual core
		Memory:    512 * models.MB,
		PidsLimit: &pidsLimit,
	}

	contId, err := machine.CreateNewContainer(msg.JobId, msg.ImageTag, resources)
	if err != nil {
		q.bigProblem("Unable to create job container", msg, err, jobDir)
		return nil, "", "", err
	}

	err = machine.CopyToContainer(contId, jobDir)
	if err != nil {
		q.bigProblem("Unable to copy files to job container", msg, err, jobDir)
		return nil, "", "", err
	}

	msg.ContainerId = contId
	res := q.db.Save(msg)
	if res.Error != nil {
		q.bigProblem("Unable to update job in db", msg, res.Error, jobDir)
		return nil, "", "", err
	}

	return machine, contId, filepath.Dir(jobDir), nil // get the dir above autolab subdir
}

// bigProblem job failed, Not good!
// The publicReason will be displayed to the end user, providing a user-friendly message.
// The err parameter holds the underlying error, used for debugging purposes.
// we use variadic param to make jobtar optional
func (q *JobQueue) bigProblem(publicReason string, job *models.Job, err error, jobTar ...string) {
	log.Error().Err(err).Str("reason", publicReason).Msgf("Job: %s failed", job.JobId)

	// todo maybe upload job tar to a "failed" bucket
	// 	so that it can be retrieved later for debugging

	// job will be saved by the cleanup function
	job.Status = models.Failed
	job.StatusMessage = publicReason
}

func (q *JobQueue) setJobAsCancelled(job *models.Job) {
	job.Status = models.Canceled
	job.StatusMessage = "Job was cancelled"
}

// greatSuccess Very nice!
// jobResult will contain the final job info, returned to the user
func (q *JobQueue) greatSuccess(job *models.Job, jobResult string) {
	log.Info().Msgf("Job: %s completed", job.JobId)

	// job will be saved by the cleanup function
	job.Status = models.Complete
	job.StatusMessage = jobResult
}

// cleanupJob clean up job
// sets job to success, removes the container and associated tmp job data
func (q *JobQueue) cleanupJob(msg *models.Job, client *docker.DkClient, submissionTmpFolder string) {
	log.Debug().Msgf("Cleaning up job: %s", msg.JobId)

	q.updateJobVeryNice(msg)

	err := client.RemoveContainer(msg.ContainerId, true, true)
	if err != nil {
		log.Error().Err(err).Msgf("Unable to remove container %s", msg.ContainerId)
	}

	q.dkSrv.ClientManager.DecreaseJobCount(msg.MachineId)

	err = os.RemoveAll(submissionTmpFolder)
	if err != nil {
		log.Error().Err(err).Msgf("Unable to remove tmp job directory %s", filepath.Dir(submissionTmpFolder))
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
		log.Error().Err(res.Error).Msgf("An error occured while saving job to db")
		// maybe fail ??
		// q.bigProblem()
	}
}

func (q *JobQueue) verifyLogs(file string, msg *models.Job) {
	if msg.Status == models.Failed {
		log.Warn().Msgf("Job %s failed skipping parsing log file", msg.JobId)
		return
	}

	outputFile, err := os.Open(file)
	if err != nil {
		q.bigProblem("Unable to open log file", msg, err, file)
		return
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Error().Err(err).Msgf("An error occured while closing log file")
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
