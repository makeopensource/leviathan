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

type Job struct {
	gorm.Model
	JobId                     string
	MachineId                 string
	ContainerId               string
	ImageTag                  string
	Status                    JobStatus
	StatusMessage             string
	StudentSubmissionFileName string
	StudentSubmissionFile     []byte
	LabData                   LabModel `gorm:"-"` // This field will be ignored by GORM
}
