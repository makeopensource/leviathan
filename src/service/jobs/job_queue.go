package jobs

import (
	"context"
	"fmt"
	cond "github.com/docker/docker/api/types/container"
	cont "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/utils"
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
}

func NewJobQueue(totalJobs uint, db *gorm.DB, dk *docker.DkService) *JobQueue {
	queue := &JobQueue{
		jobChannel: make(chan *models.Job, totalJobs),
		totalJobs:  totalJobs,
		db:         db,
		dkSrv:      dk,
	}

	queue.CreateJobProcessors()
	return queue
}

func (q *JobQueue) CreateJobProcessors() {
	for i := 1; i < int(q.totalJobs); i++ {
		go q.messageProcessors(i)
	}
}

func (q *JobQueue) messageProcessors(workerId int) {
	for msg := range q.jobChannel {
		log.Info().Msgf("Worker: %d is now processing job: %s", workerId, msg.JobId)
		q.runJob(msg)
	}
}

func (q *JobQueue) runJob(msg *models.Job) {
	q.setJobInProgress(msg)

	client, contId, jobTar := q.setupJob(msg)
	if client == nil || contId == "" {
		return
	}

	defer func() {
		q.cleanupJob(msg, client, jobTar)
	}()

	err := client.StartContainer(contId)
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
	}
}

func (q *JobQueue) AddJob(mes *models.Job) {
	q.jobChannel <- mes
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
func (q *JobQueue) setupJob(msg *models.Job) (*docker.DkClient, string, string) {
	jobTar, err := utils.ArchiveJobData(map[string][]byte{
		msg.LabData.GraderFilename:    msg.LabData.GraderFile,
		msg.LabData.MakeFilename:      msg.LabData.MakeFile,
		msg.StudentSubmissionFileName: msg.StudentSubmissionFile,
	})
	if err != nil {
		q.bigProblem("Failed to convert job data to tar", msg, err, jobTar)
		return nil, "", ""
	}

	machine, err := q.dkSrv.ClientManager.GetClientById(msg.MachineId)
	if err != nil {
		q.bigProblem("Failed to get machine info", msg, err, jobTar)
		return nil, "", ""
	}

	// todo load from job message
	resources := cont.Resources{
		Memory:   512 * 1000000,
		NanoCPUs: 2 * 1000000000,
		//PidsLimit: 50 todo
	}

	contId, err := machine.CreateNewContainer(msg.JobId, msg.ImageTag, resources, jobTar)
	if err != nil {
		q.bigProblem("Unable to create job container", msg, err, jobTar)
		return nil, "", ""
	}

	//err = machine.CopyToContainer(contId, jobTar)
	//if err != nil {
	//	q.bigProblem("Unable to copy files to job container", msg, err, jobTar)
	//	return nil, ""
	//}

	msg.ContainerId = contId
	res := q.db.Save(msg)
	if res.Error != nil {
		q.bigProblem("Unable to update job in db", msg, res.Error, jobTar)
		return nil, "", ""
	}

	return machine, contId, jobTar
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
func (q *JobQueue) cleanupJob(msg *models.Job, client *docker.DkClient, tar string) {
	q.updateJobVeryNice(msg)

	err := client.RemoveContainer(msg.ContainerId, true, true)
	if err != nil {
		log.Error().Err(err).Msgf("Unable to remove container %s", msg.ContainerId)
	}

	q.dkSrv.ClientManager.DecreaseJobCount(msg.MachineId)

	err = os.RemoveAll(filepath.Dir(tar))
	if err != nil {
		log.Error().Err(err).Msgf("Unable to remove tmp job directory %s", filepath.Dir(tar))
		return
	}
}

// Job is in progress, success soon!
func (q *JobQueue) setJobInProgress(msg *models.Job) {
	msg.Status = models.Running
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

	line, err := utils.GetLastLine(outputFile)
	if err != nil {
		q.bigProblem("unable to get logs", msg, err)
		return
	}

	if !utils.IsValidJSON(line) {
		q.bigProblem("unable to parse log output", msg, err)
		return
	}

	q.greatSuccess(msg, line)
}
