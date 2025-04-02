package jobs

import (
	"context"
	"fmt"
	dk "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/makeopensource/leviathan/common"
	md "github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/labs"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

type JobQueue struct {
	jobSemaphore *md.WorkerSemaphore
	db           *gorm.DB
	dkSrv        *docker.DkService
	contextMap   *md.Map[string, func()]
	labSrv       *labs.LabService
}

func NewJobQueue(totalJobs uint, db *gorm.DB, dk *docker.DkService, lab *labs.LabService) *JobQueue {
	queue := &JobQueue{
		contextMap:   &md.Map[string, func()]{},
		db:           db,
		dkSrv:        dk,
		jobSemaphore: md.NewWorkerSemaphore(int(totalJobs)),
		labSrv:       lab,
	}

	return queue
}

func (q *JobQueue) AddJob(mes *md.Job) error {
	jog(mes.JobCtx).Info().Msg("sending job to queue")
	err := mes.ValidateForQueue()
	if err != nil {
		return common.ErrLog("job validation failed: "+err.Error(), err, jog(mes.JobCtx).Error())
	}

	go q.worker(mes)
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

func (q *JobQueue) worker(msg *md.Job) {
	if msg == nil {
		log.Error().Msg("job received was nil, THIS SHOULD NEVER HAPPEN")
		return
	}

	select {
	case <-q.jobSemaphore.Acquire():
		defer q.jobSemaphore.Release()
		jog(msg.JobCtx).Info().Msgf("worker is now processing job")
		q.runJob(msg)
		return
	case <-msg.JobCtx.Done():
		q.setJobAsCancelled(msg)
		q.cleanupJob(msg, nil)
		jog(msg.JobCtx).Warn().Msgf("job context was canceled before queue could process")
		return
	}
}

// runJob should ALWAYS BE BLOCKING, as it prevents the worker from moving on to a new job
func (q *JobQueue) runJob(job *md.Job) {
	client, contId, err := q.setupJob(job)
	defer q.cleanupJob(job, client)
	if err != nil {
		q.bigProblem(job, err)
		return
	}

	logStatusCh := make(chan md.JobError)

	q.setJobInProgress(job)
	err2 := client.StartContainer(contId)
	if err2 != nil {
		q.bigProblem(job, md.JError("unable to start job container", err2))
		return
	}

	// start writing to log file so that,
	// we can stream changes to the log file to the user
	go func() {
		logStatusCh <- writeLogs(client, job)
	}()

	statusCh, errCh := client.WaitForContainerStatusChange(contId)
	select {
	case <-statusCh:
		if mes := <-logStatusCh; mes != nil {
			q.bigProblem(job, mes)
			return
		}
		logLine, err := verifyLogs(job)
		if err != nil {
			q.bigProblem(job, err)
			return
		}
		q.greatSuccess(job, logLine)
		return
	case err := <-errCh:
		q.bigProblem(job, md.JError("error occurred while waiting for job process", err))
		return
	case <-time.After(job.LabData.JobTimeout):
		q.bigProblem(job, md.JError(timeoutErrLog(job.LabData.JobTimeout), nil))
		return
	case <-job.JobCtx.Done():
		q.setJobAsCancelled(job)
		return
	}
}

func timeoutErrLog(timeout time.Duration) string {
	return fmt.Sprintf("Maximum timeout reached for job, job ran for %s", timeout)
}

// setupJob Set up job like king, yes!
// returns nil client if an error occurred while setup,
// make sure to handle null ptr dereference
func (q *JobQueue) setupJob(msg *md.Job) (*docker.DkClient, string, md.JobError) {
	q.setJobInSetup(msg)

	machine, err := q.dkSrv.ClientManager.GetClientById(msg.MachineId)
	if err != nil {
		return nil, "", md.JError("Failed to get machine info", err)
	}

	// incase dockerfile is not passed and referenced via tag name
	if msg.LabData.DockerFilePath != "" {
		err = machine.BuildImageFromDockerfile(msg.LabData.DockerFilePath, msg.LabData.ImageTag)
		if err != nil {
			return nil, "", md.JError("Failed to create image", err)
		}
	}

	resources := dk.Resources{
		NanoCPUs:  msg.LabData.JobLimits.NanoCPU * md.CPUQuota,
		Memory:    msg.LabData.JobLimits.Memory * md.MB,
		PidsLimit: &msg.LabData.JobLimits.PidsLimit,
	}

	contId, err := machine.CreateNewContainer(
		msg.JobId,
		msg.LabData.ImageTag,
		msg.LabData.JobEntryCmd,
		resources,
	)
	if err != nil {
		return nil, "", md.JError("unable to create job container", err)
	}

	err = machine.CopyToContainer(contId, msg.TmpJobFolderPath)
	if err != nil {
		return nil, "", md.JError("unable to copy files to job container", err)
	}

	msg.ContainerId = contId
	res := q.db.Save(msg)
	if res.Error != nil {
		return nil, "", md.JError("unable to update job in db", err)
	}

	return machine, contId, nil
}

// cleanupJob clean up job,
// updates job in DB, removes the container and associated tmp job data
func (q *JobQueue) cleanupJob(msg *md.Job, client *docker.DkClient) {
	jog(msg.JobCtx).Info().Msg("cleaning up job")
	q.updateJobVeryNice(msg)

	if client != nil {
		if err := client.RemoveContainer(msg.ContainerId, true, true); err != nil {
			jog(msg.JobCtx).Warn().Err(err).Str("containerID", msg.ContainerId).Msg("unable to remove container")
		}
	}

	q.dkSrv.ClientManager.DecreaseJobCount(msg.MachineId)
	q.contextMap.Delete(msg.JobId)

	tmpFold := filepath.Dir(msg.TmpJobFolderPath) // get the dir above autolab subdir
	if err := os.RemoveAll(tmpFold); err != nil {
		jog(msg.JobCtx).Warn().Err(err).Str("dir", tmpFold).Msg("unable to remove tmp job directory")
		return
	}
}

// greatSuccess set job status to models.Complete
//
// Very nice!
//
// jobResult is the last line expected to be valid json string, returned to the job caller
func (q *JobQueue) greatSuccess(job *md.Job, jobResult string) {
	jog(job.JobCtx).Info().Msg("job completed successfully")
	job.Status = md.Complete
	job.StatusMessage = jobResult
}

// bigProblem set job status to models.Failed
//
// job failed, Not good!
func (q *JobQueue) bigProblem(job *md.Job, jErr md.JobError) {
	jog(job.JobCtx).Error().Err(jErr.Err()).Str("reason", jErr.Reason()).Msg("job failed")
	job.Status = md.Failed
	job.StatusMessage = jErr.Reason()
	job.Error = jErr.ErrStr()
}

func (q *JobQueue) setJobAsCancelled(job *md.Job) {
	jog(job.JobCtx).Info().Msg("job was cancelled")
	job.Status = md.Canceled
	job.StatusMessage = "Job was cancelled"
}

// setJobInProgress set job status as models.Running
//
// Job is in progress, success soon!
func (q *JobQueue) setJobInProgress(msg *md.Job) {
	msg.Status = md.Running
	q.updateJobVeryNice(msg)
}

// setJobInSetup set job status as models.Preparing
//
// job is being setup standby
func (q *JobQueue) setJobInSetup(msg *md.Job) {
	msg.Status = md.Preparing
	q.updateJobVeryNice(msg)
}

// updateJobVeryNice Database updated, fresh like new wife!
func (q *JobQueue) updateJobVeryNice(msg *md.Job) {
	res := q.db.Save(msg)
	if res.Error != nil {
		jog(msg.JobCtx).Error().Err(res.Error).Msg("error occurred while saving job to db")
	}
}

func writeLogs(client *docker.DkClient, msg *md.Job) md.JobError {
	outputFile, err := os.OpenFile(msg.OutputLogFilePath, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return md.JError("unable to open log file", err)
	}

	defer func() {
		err := outputFile.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error while closing output file")
		}
	}()

	logs, err := client.TailContainerLogs(context.Background(), msg.ContainerId)
	if err != nil {
		return md.JError("unable to tail job container", err)
	}

	_, err = stdcopy.StdCopy(outputFile, outputFile, logs)
	if err != nil {
		return md.JError("unable to write to log file", err)
	}
	return nil
}

func verifyLogs(msg *md.Job) (string, md.JobError) {
	if msg.Status == md.Failed {
		return "", md.JError("Job failed, skipping parsing log file", nil)
	}

	outputFile, err := os.Open(msg.OutputLogFilePath)
	if err != nil {
		return "", md.JError("unable to open log file", err)
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			jog(msg.JobCtx).Warn().Err(err).Msg("An error occurred while closing log file")
		}
	}(outputFile)

	line, err := GetLastLine(outputFile)
	if err != nil {
		return "", md.JError("unable to get logs", err)
	}
	if !IsValidJSON(line) {
		return "", md.JError("unable to parse log output", err)
	}

	return line, nil
}
