package v1

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
	types "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/jobs"
	"strings"
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
	entryCmd := req.Msg.GetEntryCmd()

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

	if tag == "" {
		return nil, fmt.Errorf("docker image tag is empty")
	}
	if entryCmd == "" {
		return nil, fmt.Errorf("entry cmd is empty")
	}

	newJob := models.Job{
		JobEntryCmd: entryCmd,
		LabData:     models.Lab{ImageTag: strings.ToLower(strings.TrimSpace(tag))},
		JobTimeout:  time.Second * time.Duration(req.Msg.JobTimeoutInSeconds),
		JobLimits: models.MachineLimits{
			PidsLimit: int64(req.Msg.GetLimits().PidLimit),
			NanoCPU:   int64(req.Msg.GetLimits().CPUCores),
			Memory:    int64(req.Msg.GetLimits().MemoryInMb),
		},
	}

	jobId, err := job.Service.NewJob(&newJob, makeF, grader, stu, dockerfile)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.NewJobResponse{JobId: jobId})
	return res, nil
}

func (job *JobServer) StreamStatus(ctx context.Context, req *connect.Request[v1.JobLogRequest], stream *connect.ServerStream[v1.JobLogsResponse]) error {
	err := job.Service.StreamJobAndLogs(ctx, req.Msg.GetJobId(), stream)
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
