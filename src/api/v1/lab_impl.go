package v1

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/labs/v1"
	typesv1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/labs"
	"time"
)

type LabServer struct {
	Srv *labs.LabService
}

func (l LabServer) NewLab(ctx context.Context, req *connect.Request[v1.NewLabRequest]) (*connect.Response[v1.NewLabResponse], error) {
	if req.Msg.LabData.Labname == "" {
		return nil, fmt.Errorf("lab name is required")
	} else if req.Msg.LabData.JobTimeoutInSeconds == 0 {
		return nil, fmt.Errorf("0 seconds as time limit is not allowed")
	} else if req.Msg.TmpFolderId == "" {
		return nil, fmt.Errorf("tmp folder id is required")
	}

	lab := &models.Lab{
		Name:       req.Msg.LabData.Labname,
		JobTimeout: time.Duration(req.Msg.LabData.JobTimeoutInSeconds) * time.Second,
		JobLimits: models.MachineLimits{
			PidsLimit: int64(req.Msg.LabData.Limits.PidLimit),
			NanoCPU:   int64(req.Msg.LabData.Limits.CPUCores),
			Memory:    int64(req.Msg.LabData.Limits.MemoryInMb),
		},
		JobEntryCmd:       req.Msg.LabData.EntryCmd,
		AutolabCompatible: req.Msg.LabData.AutolabCompatibilityMode,
	}

	labID, err := l.Srv.CreateLab(lab, req.Msg.TmpFolderId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.NewLabResponse{LabId: int64(labID)})
	return res, nil
}

func (l LabServer) EditLab(ctx context.Context, c *connect.Request[typesv1.LabData]) (*connect.Response[v1.EditLabResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (l LabServer) DeleteLab(ctx context.Context, req *connect.Request[v1.DeleteLabRequest]) (*connect.Response[v1.DeleteLabResponse], error) {

	return nil, fmt.Errorf("unimplmented methods")
}
