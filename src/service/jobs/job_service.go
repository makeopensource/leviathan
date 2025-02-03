package jobs

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type JobService struct {
	db           *gorm.DB
	labFileCache *utils.LabFilesCache
}

func NewJobService(db *gorm.DB, cache *utils.LabFilesCache) *JobService {
	return &JobService{
		db:           db,
		labFileCache: cache,
	}
}

func (jobSrv *JobService) NewJob(imageTag string, courseName string, studentFileTarFile []byte) (string, error) {
	jobId, err := uuid.NewUUID()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to generate job ID")
		return "", fmt.Errorf("failed to generate job ID")
	}

	message := &models.JobMessage{
		JobId:          jobId.String(),
		Status:         models.Queued,
		StudentTarFile: studentFileTarFile,
		LabName:        courseName,
	}

	res := jobSrv.db.Create(message)
	if res.Error != nil {
		log.Error().Err(res.Error).Msgf("Failed to save job to db")
		return "", fmt.Errorf("failed to save job to db")
	}

	// run in go routine in case queue is full and this gets blocked
	go AddJob(message)

	return jobId.String(), nil
}

func (jobSrv *JobService) GetJobStatus(jobUuid string) error {
	return nil
}

func (jobSrv *JobService) CancelJob(jobUuid string) error {
	return nil
}
