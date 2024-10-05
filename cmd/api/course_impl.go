package api

import (
	"connectrpc.com/connect"
	"context"
	dkrpc "github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1"
	"log"
)

type DockerServer struct{}

func (dk *DockerServer) CreateContainer(ctx context.Context, req *connect.Request[dkrpc.CreateContainerRequest]) (*connect.Response[dkrpc.CreateContainerResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&dkrpc.CreateContainerResponse{})
	return res, nil
}

func (dk *DockerServer) DeleteContainer(ctx context.Context, req *connect.Request[dkrpc.DeleteContainerRequest]) (*connect.Response[dkrpc.DeleteContainerResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&dkrpc.DeleteContainerResponse{})
	return res, nil
}

func (dk *DockerServer) ListContainers(ctx context.Context, req *connect.Request[dkrpc.ListContainersRequest]) (*connect.Response[dkrpc.ListContainersResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&dkrpc.ListContainersResponse{})
	return res, nil
}

func (dk *DockerServer) Echo(ctx context.Context, req *connect.Request[dkrpc.EchoRequest]) (*connect.Response[dkrpc.EchoResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&dkrpc.EchoResponse{})
	return res, nil
}
