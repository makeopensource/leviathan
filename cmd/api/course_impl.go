package api

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/docker/docker/client"
	"github.com/makeopensource/leviathan/internal/dockerclient"
	dkrpc "github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1"
	"github.com/rs/zerolog/log"
)

type DockerServer struct {
	clientList []*client.Client
}

func (dk *DockerServer) CreateContainer(ctx context.Context, req *connect.Request[dkrpc.CreateContainerRequest]) (*connect.Response[dkrpc.CreateContainerResponse], error) {
	res := connect.NewResponse(&dkrpc.CreateContainerResponse{})
	return res, nil
}

func (dk *DockerServer) DeleteContainer(ctx context.Context, req *connect.Request[dkrpc.DeleteContainerRequest]) (*connect.Response[dkrpc.DeleteContainerResponse], error) {
	containerId := req.Msg.GetContainerId()
	log.Debug().Str("Container ID", req.Msg.GetContainerId()).Msgf("Recivied delete container request")

	if containerId == "" {
		return nil, errors.New("no container Id found")
	}

	res := connect.NewResponse(&dkrpc.DeleteContainerResponse{})
	return res, nil
}

func (dk *DockerServer) ListContainers(ctx context.Context, req *connect.Request[dkrpc.ListContainersRequest]) (*connect.Response[dkrpc.ListContainersResponse], error) {
	res := connect.NewResponse(&dkrpc.ListContainersResponse{})
	return res, nil
}

func (dk *DockerServer) StartContainer(ctx context.Context, req *connect.Request[dkrpc.StartContainerRequest]) (*connect.Response[dkrpc.StartContainerResponse], error) {
	res := connect.NewResponse(&dkrpc.StartContainerResponse{})
	return res, nil
}
func (dk *DockerServer) StopContainer(ctx context.Context, req *connect.Request[dkrpc.StopContainerRequest]) (*connect.Response[dkrpc.StopContainerResponse], error) {
	res := connect.NewResponse(&dkrpc.StopContainerResponse{})
	return res, nil
}
func (dk *DockerServer) CreateNewImage(ctx context.Context, req *connect.Request[dkrpc.NewImageRequest]) (*connect.Response[dkrpc.NewImageResponse], error) {
	filename := req.Msg.File.GetFilename()
	contents := req.Msg.File.GetContent()
	imageTag := req.Msg.GetImageTag()

	err := dockerclient.HandleNewImageReq(filename, contents, imageTag, dk.clientList)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&dkrpc.NewImageResponse{})

	return res, nil
}
func (dk *DockerServer) ListImages(ctx context.Context, req *connect.Request[dkrpc.ListImageRequest]) (*connect.Response[dkrpc.ListImageResponse], error) {
	images := dockerclient.HandleListImagesReq(dk.clientList)
	res := connect.NewResponse(&dkrpc.ListImageResponse{Images: images})
	return res, nil
}
