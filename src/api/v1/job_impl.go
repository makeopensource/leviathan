package v1

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/jobs/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/jobs"
)

type JobServer struct {
	Service *jobs.JobService
}

func (job *JobServer) NewJob(ctx context.Context, req *connect.Request[v1.NewJobRequest]) (*connect.Response[v1.NewJobResponse], error) {
	makeF := req.Msg.GetMakeFile()
	grader := req.Msg.GetGraderFile()
	stu := req.Msg.GetStudentSubmission()
	tag := req.Msg.GetImageTag()

	newJob := models.Job{
		ImageTag:                  tag,
		StudentSubmissionFileName: stu.Filename,
		StudentSubmissionFile:     stu.Content,
		LabData: models.LabModel{
			GraderFilename: grader.Filename,
			GraderFile:     grader.Content,
			MakeFilename:   makeF.Filename,
			MakeFile:       makeF.Content,
		},
	}

	jobId, err := job.Service.NewJob(&newJob)
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
