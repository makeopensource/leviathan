package routes

import (
	"context"
	"github.com/UBAutograding/leviathan/internal/rpc/V1/jobs"
)

// grpc implementation for jobs service

type JobsServiceSrv struct {
	jobs.UnimplementedJobServiceServer
}

func (d *JobsServiceSrv) NewJob(_ context.Context, request *jobs.NewJobRequest) (*jobs.NewJobResponse, error) {
	return &jobs.NewJobResponse{}, nil
}

func (d *JobsServiceSrv) DeleteContainer(_ context.Context, request *jobs.JobStatusRequest) (*jobs.JobStatusResponse, error) {
	return &jobs.JobStatusResponse{}, nil
}

func (d *JobsServiceSrv) CancelJob(_ context.Context, request *jobs.CancelJobRequest) (*jobs.CancelJobResponse, error) {
	return &jobs.CancelJobResponse{}, nil
}

func (d *JobsServiceSrv) Echo(_ context.Context, request *jobs.EchoRequest) (*jobs.EchoResponse, error) {
	return &jobs.EchoResponse{MessageResponse: request.Message}, nil
}
