package jobs

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

type JobService struct {
	db           *gorm.DB
	labFileCache *utils.LabFilesCache
	dockerSrv    *docker.DockerService
	queue        *JobQueue
}

func NewJobService(db *gorm.DB, cache *utils.LabFilesCache, dockerService *docker.DockerService) *JobService {
	return &JobService{
		db:           db,
		labFileCache: cache,
		dockerSrv:    dockerService,
		queue:        NewJobQueue(30, db, dockerService),
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

	// run in go routine in case queue is full and this gets blocked
	go job.queue.AddJob(jobReq)

	return jobId.String(), nil
}

func (job *JobService) GetJobStatus(jobUuid string) error {
	return nil
}

// WaitForJob is similar to GetJobStatus but is blocking and runs until job is complete or errors
func (job *JobService) WaitForJob(jobUuid string) (*models.Job, error) {
	// keep checking job until complete
	for {
		time.Sleep(5 * time.Second)

		info, err := job.getJob(jobUuid)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to get job")
			continue
		}

		if info.Status == models.Complete || info.Status == models.Failed {
			return info, nil
		}
	}
}

func (job *JobService) getJob(jobUuid string) (*models.Job, error) {
	var jobInfo *models.Job
	res := job.db.Find(&jobInfo, "job_id = ?", jobUuid)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get job info from db")
	}
	return jobInfo, nil
}

func (job *JobService) CancelJob(jobUuid string) error {
	return nil
}

// setupLogFile store grader output
// this is blocking operation make sure to
// stream logs in a go routine
func (job *JobService) setupLogFile(msg *models.Job) *os.File {
	outputFile := fmt.Sprintf("%s/%s.txt", utils.OutputFolder.GetStr(), msg.JobId)
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
