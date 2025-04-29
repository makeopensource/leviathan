package labs

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	labrpc "github.com/makeopensource/leviathan/generated/labs/v1"
	rpctypes "github.com/makeopensource/leviathan/generated/types/v1"
	"time"
)

type LabServer struct {
	Srv *LabService
}

func (l LabServer) NewLab(ctx context.Context, req *connect.Request[labrpc.NewLabRequest]) (*connect.Response[labrpc.NewLabResponse], error) {
	if req.Msg.LabData.Labname == "" {
		return nil, fmt.Errorf("lab name is required")
	} else if req.Msg.LabData.JobTimeoutInSeconds == 0 {
		return nil, fmt.Errorf("0 seconds as time limit is not allowed")
	} else if req.Msg.TmpFolderId == "" {
		return nil, fmt.Errorf("tmp folder id is required")
	}

	lab := &Lab{
		Name:       req.Msg.LabData.Labname,
		JobTimeout: time.Duration(req.Msg.LabData.JobTimeoutInSeconds) * time.Second,
		JobLimits: MachineLimits{
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

	res := connect.NewResponse(&labrpc.NewLabResponse{LabId: int64(labID)})
	return res, nil
}

func (l LabServer) EditLab(ctx context.Context, c *connect.Request[rpctypes.LabData]) (*connect.Response[labrpc.EditLabResponse], error) {
	// todo
	return nil, fmt.Errorf("unimplmented")
}

func (l LabServer) DeleteLab(ctx context.Context, req *connect.Request[labrpc.DeleteLabRequest]) (*connect.Response[labrpc.DeleteLabResponse], error) {
	// todo
	return nil, fmt.Errorf("unimplmented")
}
