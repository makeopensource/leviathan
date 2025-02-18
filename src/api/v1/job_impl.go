package v1

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
	types "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/jobs"
	"time"
)

type JobServer struct {
	Service *jobs.JobService
}

func (job *JobServer) NewJob(ctx context.Context, req *connect.Request[v1.NewJobRequest]) (*connect.Response[v1.NewJobResponse], error) {
	makeF := req.Msg.GetMakeFile()
	grader := req.Msg.GetGraderFile()
	stu := req.Msg.GetStudentSubmission()
	tag := req.Msg.GetImageName()
	dockerfile := req.Msg.GetDockerFile()

	vars := []*types.FileUpload{makeF, grader, stu, dockerfile}
	for _, v := range vars {
		if v == nil {
			return nil, fmt.Errorf("some fields are empty")
		} else if v.Filename == "" {
			return nil, fmt.Errorf("filename is empty")
		} else if len(v.Content) == 0 {
			return nil, fmt.Errorf("empty content for file %s", v.Filename)
		}
	}

	newJob := models.Job{
		LabData:    models.LabModel{ImageTag: tag},
		JobTimeout: time.Second * time.Duration(req.Msg.JobTimeoutInSeconds),
	}

	jobId, err := job.Service.NewJob(&newJob, makeF, grader, stu, dockerfile)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.NewJobResponse{JobId: jobId})
	return res, nil
}

func (job *JobServer) StreamJobLogs(ctx context.Context, req *connect.Request[v1.JobLogRequest], responseStream *connect.ServerStream[v1.JobLogsResponse]) error {
	_, _, err := job.Service.ListenToJobLogs(ctx, req.Msg.GetJobId())
	if err != nil {
		return err
	}

	return nil
}

func (job *JobServer) CancelJob(ctx context.Context, req *connect.Request[v1.CancelJobRequest]) (*connect.Response[v1.CancelJobResponse], error) {
	msgId := req.Msg.GetJobId()
	if msgId == "" {
		return nil, fmt.Errorf("job Id is empty")
	}

	job.Service.CancelJob(msgId)
	return connect.NewResponse(&v1.CancelJobResponse{}), nil
}
