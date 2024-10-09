package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/docker/docker/client"
	v1 "github.com/makeopensource/leviathan/internal/generated/jobs/v1"
)

type JobServer struct {
	clientList map[string]*client.Client
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
