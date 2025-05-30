package jobs

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
)

type JobServer struct {
	srv *JobService
}

func NewJobServer(srv *JobService) *JobServer {
	return &JobServer{srv: srv}
}

func (job *JobServer) NewJob(_ context.Context, req *connect.Request[v1.NewJobRequest]) (*connect.Response[v1.NewJobResponse], error) {
	labId := req.Msg.LabID
	submissionTmpFolder := req.Msg.TmpSubmissionFolderId
	if submissionTmpFolder == "" {
		return nil, fmt.Errorf("submission folder id is empty")
	}

	newJob := &Job{LabID: uint(labId)}
	jobId, err := job.srv.NewJob(newJob, submissionTmpFolder)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.NewJobResponse{JobId: jobId})
	return res, nil
}

func (job *JobServer) GetStatus(_ context.Context, req *connect.Request[v1.JobLogRequest]) (*connect.Response[v1.JobLogsResponse], error) {
	status, logs, err := job.srv.GetJobStatusAndLogs(req.Msg.GetJobId())
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.JobLogsResponse{
		JobInfo: status.ToProto(),
		Logs:    logs,
	})
	return res, nil
}

func (job *JobServer) StreamStatus(ctx context.Context, req *connect.Request[v1.JobLogRequest], stream *connect.ServerStream[v1.JobLogsResponse]) error {
	streamFunc := func(jobInfo *Job, logs string) error {
		return stream.Send(&v1.JobLogsResponse{
			JobInfo: jobInfo.ToProto(),
			Logs:    logs,
		})
	}

	err := job.srv.StreamJobAndLogs(ctx, req.Msg.GetJobId(), streamFunc)
	if err != nil {
		return err
	}
	return nil
}

func (job *JobServer) CancelJob(_ context.Context, req *connect.Request[v1.CancelJobRequest]) (*connect.Response[v1.CancelJobResponse], error) {
	msgId := req.Msg.GetJobId()
	if msgId == "" {
		return nil, fmt.Errorf("job Id is empty")
	}

	job.srv.CancelJob(msgId)
	return connect.NewResponse(&v1.CancelJobResponse{}), nil
}
