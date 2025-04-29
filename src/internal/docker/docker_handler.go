package docker

import (
	"connectrpc.com/connect"
	"context"
	dkrpc "github.com/makeopensource/leviathan/generated/docker_rpc/v1"
)

type Server struct {
	Service *DkService
}

func (dk *Server) CreateContainer(_ context.Context, req *connect.Request[dkrpc.CreateContainerRequest]) (*connect.Response[dkrpc.CreateContainerResponse], error) {
	res := connect.NewResponse(&dkrpc.CreateContainerResponse{})
	return res, nil
}

func (dk *Server) StartContainer(_ context.Context, req *connect.Request[dkrpc.StartContainerRequest]) (*connect.Response[dkrpc.StartContainerResponse], error) {
	res := connect.NewResponse(&dkrpc.StartContainerResponse{})
	return res, nil
}

func (dk *Server) DeleteContainer(_ context.Context, req *connect.Request[dkrpc.DeleteContainerRequest]) (*connect.Response[dkrpc.DeleteContainerResponse], error) {
	res := connect.NewResponse(&dkrpc.DeleteContainerResponse{})
	return res, nil
}

func (dk *Server) StopContainer(_ context.Context, req *connect.Request[dkrpc.StopContainerRequest]) (*connect.Response[dkrpc.StopContainerResponse], error) {
	res := connect.NewResponse(&dkrpc.StopContainerResponse{})
	return res, nil
}

func (dk *Server) GetContainerLogs(_ context.Context, req *connect.Request[dkrpc.GetContainerLogRequest], responseStream *connect.ServerStream[dkrpc.GetContainerLogResponse]) error {
	return nil
}

func (dk *Server) CreateNewImage(_ context.Context, req *connect.Request[dkrpc.NewImageRequest]) (*connect.Response[dkrpc.NewImageResponse], error) {
	res := connect.NewResponse(&dkrpc.NewImageResponse{})
	return res, nil
}
func (dk *Server) ListImages(_ context.Context, _ *connect.Request[dkrpc.ListImageRequest]) (*connect.Response[dkrpc.ListImageResponse], error) {
	res := connect.NewResponse(&dkrpc.ListImageResponse{})
	return res, nil
}

func (dk *Server) ListContainers(_ context.Context, _ *connect.Request[dkrpc.ListContainersRequest]) (*connect.Response[dkrpc.ListContainersResponse], error) {
	res := connect.NewResponse(&dkrpc.ListContainersResponse{})
	return res, nil
}
