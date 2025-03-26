package models

import (
	"context"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type JobStatus string

// Done indicate the job has been tried by the queue and processed.
// i.e: the JobStatus is either Failed, Complete, Canceled
func (js JobStatus) Done() bool {
	if js == Failed || js == Complete || js == Canceled {
		return true
	}
	return false
}

// go way of doing enums
const (
	// Queued -> job was sent to the job channel waiting to be picked by a worker
	Queued JobStatus = "queued"
	// Preparing -> job is picked up by a worker
	// and the required setup is being done.
	Preparing JobStatus = "preparing"
	// Running leviathan has successfully started the grading container
	// and is waiting for it to end
	Running JobStatus = "running"
	// Complete -> indicates job is complete and
	// leviathan was able to parse the log line correctly
	Complete JobStatus = "complete"
	// Failed -> job failed or the log line was unable to be parsed as json
	Failed JobStatus = "failed"
	// Canceled -> job was cancelled by the user
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
	JobId         string `gorm:"uniqueIndex"`
	MachineId     string
	ContainerId   string
	Status        JobStatus
	StatusMessage string
	// to store if an error occurred, otherwise empty,
	Error   string
	LabID   uint
	LabData *Lab `gorm:"foreignKey:LabID;references:ID"`
	// OutputLogFilePath text file contain the container std out
	OutputLogFilePath string
	// TmpJobFolderPath holds the path to the tmp dir all files related to the job except the final output
	TmpJobFolderPath string
	JobCtx           context.Context `gorm:"-"`
}

func (j *Job) ToProto() *v1.JobStatus {
	return &v1.JobStatus{
		JobId:         j.JobId,
		Status:        string(j.Status),
		StatusMessage: j.StatusMessage,
	}
}

// ValidateForQueue checks for fields required before sending job to queue
func (j *Job) ValidateForQueue() error {
	if j == nil {
		return fmt.Errorf("job is nil")
	}
	if j.JobId == "" {
		return fmt.Errorf("job id is empty")
	}
	if j.MachineId == "" {
		return fmt.Errorf("machine id is empty")
	}
	if j.LabData.JobEntryCmd == "" {
		return fmt.Errorf("job entry cmd is empty")
	}
	if j.LabData.JobTimeout == 0 {
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
	ch := tx.Statement.Context.Value(BroadcastKey)
	if ch == nil {
		log.Warn().Msg("database broadcast channel is nil")
		return
	}
	// always cast after checking if the value exists,
	// prevent null ptr deref
	go ch.(*BroadcastChannel).Broadcast(j)
	return
}
