package models

import (
	"gorm.io/gorm"
)

type JobStatus string

const (
	Queued   JobStatus = "queued"
	Running  JobStatus = "running"
	Complete JobStatus = "complete"
	Failed   JobStatus = "failed"
)

type JobMessage struct {
	gorm.Model
	JobId          string
	Status         JobStatus
	StudentTarFile []byte
	LabName        string
	LabData        LabModel `gorm:"-"` // This field will be ignored by GORM
}
