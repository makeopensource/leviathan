package api

import (
	"connectrpc.com/connect"
	"context"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
	"github.com/makeopensource/leviathan/service/jobs"
)

type JobServer struct {
	service *jobs.JobService
}

func (job *JobServer) NewJob(ctx context.Context, req *connect.Request[v1.NewJobRequest]) (*connect.Response[v1.NewJobResponse], error) {
	res := connect.NewResponse(&v1.NewJobResponse{})
	return res, nil
}
func (job *JobServer) JobStatus(ctx context.Context, req *connect.Request[v1.JobStatusRequest]) (*connect.Response[v1.JobStatusResponse], error) {
	res := connect.NewResponse(&v1.JobStatusResponse{})
	return res, nil
}

func (job *JobServer) CancelJob(ctx context.Context, req *connect.Request[v1.CancelJobRequest]) (*connect.Response[v1.CancelJobResponse], error) {
	res := connect.NewResponse(&v1.CancelJobResponse{})
	return res, nil
}
