package database

import (
	"github.com/makeopensource/leviathan/internal/jobs"
	"gorm.io/gorm"
)

type Service struct {
	*gorm.DB
	*JobDatabase
	*LabDatabase
}

// NewDatabaseWithGorm calls initDB implicitly
func NewDatabaseWithGorm() (*Service, *jobs.BroadcastChannel) {
	db, bc := initDB()
	return NewDatabase(db), bc
}

func NewDatabase(db *gorm.DB) *Service {
	return &Service{
		DB:          db,
		JobDatabase: &JobDatabase{db: db},
		LabDatabase: &LabDatabase{db: db},
	}
}
