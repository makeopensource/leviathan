package routes

import (
	"context"
	"github.com/UBAutograding/leviathan/internal/rpc/V1/docker"
)

// grpc implementation for docker

type DockerServiceSrv struct {
	docker.UnimplementedDockerServiceServer
}

func (d *DockerServiceSrv) CreateContainer(_ context.Context, request *docker.CreateRequest) (*docker.CreateResponse, error) {
	return &docker.CreateResponse{}, nil
}

func (d *DockerServiceSrv) DeleteContainer(_ context.Context, request *docker.DeleteRequest) (*docker.DeleteResponse, error) {
	return &docker.DeleteResponse{}, nil
}
func (d *DockerServiceSrv) ListContainers(_ context.Context, request *docker.ListRequest) (*docker.ListResponse, error) {
	return &docker.ListResponse{}, nil
}
func (d *DockerServiceSrv) Echo(_ context.Context, request *docker.EchoRequest) (*docker.EchoResponse, error) {
	return &docker.EchoResponse{}, nil
}
