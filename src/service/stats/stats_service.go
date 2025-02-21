package stats

import (
	"gorm.io/gorm"
)

// endpoint for other misc commands as needed

type StatService struct {
	db *gorm.DB
}

func NewStatsService(db *gorm.DB) *StatService {
	return &StatService{
		db: db,
	}
}
