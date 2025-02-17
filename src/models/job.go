package models

import (
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

type JobStatus string

// job status enum
const (
	Queued   JobStatus = "queued"
	Running  JobStatus = "running"
	Complete JobStatus = "complete"
	Failed   JobStatus = "failed"
	Canceled JobStatus = "canceled"
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
	JobId                     string `gorm:"uniqueIndex"`
	MachineId                 string
	ContainerId               string
	ImageTag                  string
	Status                    JobStatus
	StatusMessage             string
	StudentSubmissionFileName string
	StudentSubmissionFile     []byte   `gorm:"-"`
	LabData                   LabModel `gorm:"-"`
	//JobLimits                 MachineLimits
	OutputFilePath string
	JobTimeout     time.Duration
	JobCtx         context.Context `gorm:"-"`
}

// AfterUpdate adds hooks for job streaming, updates a go channel everytime a job is updated
// the consumer is responsible if it wants to use the job
func (j *Job) AfterUpdate(tx *gorm.DB) (err error) {
	ch := tx.Statement.Context.Value("broadcast").(*BroadcastChannel)
	if ch == nil {
		log.Warn().Msg("database broadcast channel is nil")
		return
	}
	go ch.Broadcast(j)
	return
}

type MachineLimits struct {
	PidsLimit int64
	// NanoCPU will be multiplied by CPUQuota
	NanoCPU uint64
	// Memory in MB
	Memory uint64
}

// LogChannelWriter implements io.Writer interface,
// to send docker output to a grpc stream
type LogChannelWriter struct {
	Channel chan string
}

func (w *LogChannelWriter) Write(p []byte) (n int, err error) {
	w.Channel <- string(p)
	return len(p), nil
}
