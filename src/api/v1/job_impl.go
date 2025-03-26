package v1

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/jobs"
	"strings"
	"time"
)

type JobServer struct {
	srv *jobs.JobService
}

func NewJobServer(srv *jobs.JobService) *JobServer {
	return &JobServer{srv: srv}
}

func (job *JobServer) NewJob(ctx context.Context, req *connect.Request[v1.NewJobRequest]) (*connect.Response[v1.NewJobResponse], error) {
	jobFiles := req.Msg.GetJobFiles()
	tag := req.Msg.GetImageName()
	dockerfile := req.Msg.GetDockerFile()
	entryCmd := req.Msg.GetEntryCmd()

	for i, v := range jobFiles {
		if v == nil {
			return nil, fmt.Errorf("nil file at index %d", i)
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

	lab := models.Lab{
		ImageTag:    strings.ToLower(strings.TrimSpace(tag)),
		JobEntryCmd: entryCmd,
		JobTimeout:  time.Second * time.Duration(req.Msg.JobTimeoutInSeconds),
		JobLimits: models.MachineLimits{
			PidsLimit: int64(req.Msg.GetLimits().PidLimit),
			NanoCPU:   int64(req.Msg.GetLimits().CPUCores),
			Memory:    int64(req.Msg.GetLimits().MemoryInMb),
		},
	}
	newJob := &models.Job{LabData: &lab}

	jobId, err := job.srv.NewJobFromRPC(newJob, jobFiles, dockerfile)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.NewJobResponse{JobId: jobId})
	return res, nil
}

func (job *JobServer) StreamStatus(ctx context.Context, req *connect.Request[v1.JobLogRequest], stream *connect.ServerStream[v1.JobLogsResponse]) error {
	streamFunc := func(jobInfo *models.Job, logs string) error {
		return stream.Send(&v1.JobLogsResponse{
			JobInfo: &v1.JobStatus{
				JobId:         jobInfo.JobId,
				Status:        string(jobInfo.Status),
				StatusMessage: jobInfo.StatusMessage,
			},
			Logs: logs,
		})
	}

	err := job.srv.StreamJobAndLogs(ctx, req.Msg.GetJobId(), streamFunc)
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

	job.srv.CancelJob(msgId)
	return connect.NewResponse(&v1.CancelJobResponse{}), nil
}
