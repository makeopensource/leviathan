package jobs

// JobStore contains the database interface needed by jobs
type JobStore interface {
	CreateJob(job *Job) error
	GetJobByUuid(jobUuid string) (*Job, error)
	UpdateJob(job *Job) error
	FetchInProgressJobs(result *[]Job) error
}
