package models

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

type JobStatus string

// Done indicate the job has been tried by the queue, and processed.
// i.e: the JobStatus is either Failed, Complete, Canceled
func (js JobStatus) Done() bool {
	if js == Failed || js == Complete || js == Canceled {
		return true
	}
	return false
}

// job status enum
const (
	Queued    JobStatus = "queued"
	Preparing JobStatus = "preparing"
	Running   JobStatus = "running"
	Complete  JobStatus = "complete"
	Failed    JobStatus = "failed"
	Canceled  JobStatus = "canceled"
)

// general resource units for docker
const (
	// CPUQuota 1 CPU in nanocores
	CPUQuota       = 1_000_000_000
	KB       int64 = 1024
	MB             = KB * 1024
)

type Job struct {
	gorm.Model
	JobId         string `gorm:"uniqueIndex"`
	MachineId     string
	ContainerId   string
	JobEntryCmd   string
	Status        JobStatus
	StatusMessage string
	LabData       LabModel `gorm:"-"`
	//JobLimits                 MachineLimits
	// OutputLogFilePath text file contain the container std out
	OutputLogFilePath string
	// TmpJobFolderPath holds the path to the tmp dir all files related to the job except the final output
	TmpJobFolderPath string
	JobTimeout       time.Duration
	JobCtx           context.Context `gorm:"-"`
}

// ValidateForQueue checks for fields required before sending job to queue
func (j *Job) ValidateForQueue() error {
	if j.JobId == "" {
		return fmt.Errorf("job id is empty")
	}
	if j.MachineId == "" {
		return fmt.Errorf("machine id is empty")
	}
	if j.JobEntryCmd == "" {
		return fmt.Errorf("job entry cmd is empty")
	}
	if j.JobTimeout == 0 {
		return fmt.Errorf("job timeout is 0")
	}
	if j.JobCtx == nil {
		return fmt.Errorf("job context is nil")
	}
	if j.OutputLogFilePath == "" {
		return fmt.Errorf("output log file is empty")
	}
	if j.TmpJobFolderPath == "" {
		return fmt.Errorf("tmp job folder is empty")
	}
	if j.LabData.ImageTag == "" {
		return fmt.Errorf("image tag is empty")
	}

	return nil
}

// AfterUpdate adds hooks for job streaming, updates a go channel everytime a job is updated
// the consumer is responsible if it wants to use the job
func (j *Job) AfterUpdate(tx *gorm.DB) (err error) {
	ch := tx.Statement.Context.Value("broadcast")
	if ch == nil {
		log.Warn().Msg("database broadcast channel is nil")
		return
	}
	// always cast after checking if the value exists,
	// prevent null ptr deref
	go ch.(*BroadcastChannel).Broadcast(j)
	return
}

type MachineLimits struct {
	PidsLimit int64
	// NanoCPU will be multiplied by CPUQuota
	NanoCPU uint64
	// Memory in MB
	Memory uint64
}
