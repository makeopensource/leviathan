package v1

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	v1 "github.com/makeopensource/leviathan/generated/labs/v1"
	"github.com/makeopensource/leviathan/service/labs"
)

type LabServer struct {
	Service *labs.LabService
}

func (labSrv *LabServer) NewLab(ctx context.Context, req *connect.Request[v1.LabRequest]) (*connect.Response[v1.NewLabResponse], error) {
	grader := req.Msg.GetGraderFile()
	makeFile := req.Msg.GetMakeFile()
	labName := req.Msg.GetLabName()

	if labName == "" {
		return nil, errors.New("empty labname")
	} else if makeFile.GetContent() == nil || len(makeFile.GetContent()) == 0 || makeFile.GetFilename() == "" {
		return nil, errors.New("empty makefile")
	} else if grader.GetContent() == nil || len(grader.GetContent()) == 0 || grader.GetFilename() == "" {
		return nil, errors.New("empty graderfile")
	}

	res := connect.NewResponse(&v1.NewLabResponse{})
	return res, nil
}
func (labSrv *LabServer) EditLab(ctx context.Context, req *connect.Request[v1.LabRequest]) (*connect.Response[v1.EditLabResponse], error) {
	res := connect.NewResponse(&v1.EditLabResponse{})
	return res, nil

}
func (labSrv *LabServer) DeleteLab(ctx context.Context, req *connect.Request[v1.DeleteLabRequest]) (*connect.Response[v1.DeleteLabResponse], error) {
	labName := req.Msg.GetLabName()

	err := labSrv.Service.DeleteLab(labName)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.DeleteLabResponse{})
	return res, nil
}
