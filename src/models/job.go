package models

import (
	"connectrpc.com/connect"
	"context"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
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
	JobLimits                 MachineLimits
	OutputFilePath            string
	JobTimeout                time.Duration
	JobCtx                    context.Context `gorm:"-"`
}

type MachineLimits struct {
	PidsLimit int64
	// NanoCPU will be multiplied by CPUQuota
	NanoCPU uint64
	// Memory in MB
	Memory uint64
}

// LogStreamWriter implements io.Writer interface,
// to send docker output to a grpc stream
type LogStreamWriter struct {
	Stream *connect.ServerStream[v1.JobLogsResponse]
}

func (w *LogStreamWriter) Write(p []byte) (n int, err error) {
	err = w.Stream.Send(&v1.JobLogsResponse{Logs: string(p)})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
