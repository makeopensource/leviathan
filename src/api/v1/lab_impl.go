package v1

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/labs/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/labs"
	"time"
)

type LabServer struct {
	Srv *labs.LabService
}

func (l LabServer) NewLab(ctx context.Context, req *connect.Request[v1.LabRequest]) (*connect.Response[v1.NewLabResponse], error) {
	if req.Msg.LabName == "" {
		return nil, fmt.Errorf("lab name is required")
	} else if req.Msg.EntryCommand == "" {
		return nil, fmt.Errorf("entry command is required")
	} else if req.Msg.TimeLimitInSeconds == 0 {
		return nil, fmt.Errorf("0 seconds as time limit is not allowed")
	}

	lab := &models.Lab{
		Name:       req.Msg.LabName,
		JobTimeout: time.Duration(req.Msg.TimeLimitInSeconds) * time.Second,
		JobLimits: models.MachineLimits{
			PidsLimit: int64(req.Msg.MachineLimits.PidLimit),
			NanoCPU:   int64(req.Msg.MachineLimits.CPUCores),
			Memory:    int64(req.Msg.MachineLimits.MemoryInMb),
		},
		JobEntryCmd: req.Msg.EntryCommand,
	}

	labID, err := l.Srv.CreateLab(lab, req.Msg.DockerFile, req.Msg.JobFiles)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.NewLabResponse{LabId: int64(labID)})
	return res, nil
}

func (l LabServer) EditLab(ctx context.Context, req *connect.Request[v1.LabRequest]) (*connect.Response[v1.EditLabResponse], error) {
	return nil, fmt.Errorf("unimplmented methods")
}

func (l LabServer) DeleteLab(ctx context.Context, req *connect.Request[v1.DeleteLabRequest]) (*connect.Response[v1.DeleteLabResponse], error) {
	return nil, fmt.Errorf("unimplmented methods")
}
