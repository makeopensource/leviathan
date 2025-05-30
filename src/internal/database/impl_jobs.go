package database

import (
	"github.com/makeopensource/leviathan/internal/jobs"
	"gorm.io/gorm"
)

// JobDatabase implements jobs.JobStore interface
type JobDatabase struct {
	db *gorm.DB
}

func (j *JobDatabase) CreateJob(job *jobs.Job) error {
	res := j.db.Create(job)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (j *JobDatabase) GetJobByUuid(jobUuid string) (*jobs.Job, error) {
	var job jobs.Job
	res := j.db.First(&job, "job_id = ?", jobUuid)
	if res.Error != nil {
		return nil, res.Error
	}
	return &job, nil
}

func (j *JobDatabase) UpdateJob(job *jobs.Job) error {
	res := j.db.Save(job)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (j *JobDatabase) FetchInProgressJobs(result *[]jobs.Job) error {
	res := j.db.Preload("LabData").
		Where("status = ?", string(jobs.Queued)).
		Or("status = ?", string(jobs.Running)).
		Or("status = ?", string(jobs.Preparing)).
		Find(result)

	if res.Error != nil {
		return res.Error
	}
	return nil
}
