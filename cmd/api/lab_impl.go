package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/docker/docker/client"
	v1 "github.com/makeopensource/leviathan/internal/generated/labs/v1"
)

type LabServer struct {
	clientList map[string]*client.Client
}

func (labSrv *LabServer) NewLab(context.Context, *connect.Request[v1.LabRequest]) (*connect.Response[v1.NewLabResponse], error) {
	res := connect.NewResponse(&v1.NewLabResponse{})
	return res, nil

}
func (labSrv *LabServer) EditLab(context.Context, *connect.Request[v1.LabRequest]) (*connect.Response[v1.EditLabResponse], error) {
	res := connect.NewResponse(&v1.EditLabResponse{})
	return res, nil

}
func (labSrv *LabServer) DeleteLab(context.Context, *connect.Request[v1.DeleteLabRequest]) (*connect.Response[v1.DeleteLabResponse], error) {
	res := connect.NewResponse(&v1.DeleteLabResponse{})
	return res, nil
}
