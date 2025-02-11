package models

import (
	"connectrpc.com/connect"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
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
	LabData                   LabModel `gorm:"-"`
	OutputFilePath            string
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
