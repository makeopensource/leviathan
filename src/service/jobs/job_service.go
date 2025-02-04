package jobs

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
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

	jobReq.JobId = jobId.String()
	jobReq.MachineId = job.dockerSrv.ClientManager.GetLeastJobCountMachineId()
	jobReq.Status = models.Queued

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

func (job *JobService) CancelJob(jobUuid string) error {
	return nil
}
