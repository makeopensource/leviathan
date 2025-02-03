package stats

import (
	"github.com/makeopensource/leviathan/utils"
	"gorm.io/gorm"
)

// endpoint for other misc commands as needed

type StatService struct {
	db            *gorm.DB
	labFilesCache *utils.LabFilesCache
}

func NewStatsService(db *gorm.DB, labFilesCache *utils.LabFilesCache) *StatService {
	return &StatService{
		db:            db,
		labFilesCache: labFilesCache,
	}
}
