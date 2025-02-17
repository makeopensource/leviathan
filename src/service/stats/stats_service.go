package stats

import (
	"github.com/makeopensource/leviathan/models"
	"gorm.io/gorm"
)

// endpoint for other misc commands as needed

type StatService struct {
	db            *gorm.DB
	labFilesCache *models.LabFilesCache
}

func NewStatsService(db *gorm.DB, labFilesCache *models.LabFilesCache) *StatService {
	return &StatService{
		db:            db,
		labFilesCache: labFilesCache,
	}
}
