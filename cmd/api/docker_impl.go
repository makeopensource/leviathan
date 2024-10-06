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

func (dk *DockerServer) CreateContainer(_ context.Context, req *connect.Request[dkrpc.CreateContainerRequest]) (*connect.Response[dkrpc.CreateContainerResponse], error) {
	res := connect.NewResponse(&dkrpc.CreateContainerResponse{})
	return res, nil
}

func (dk *DockerServer) StartContainer(_ context.Context, req *connect.Request[dkrpc.StartContainerRequest]) (*connect.Response[dkrpc.StartContainerResponse], error) {
	err := dockerclient.HandleStartContainerReq(dk.clientList, req.Msg.GetCombinedId())
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&dkrpc.StartContainerResponse{})
	return res, nil
}

func (dk *DockerServer) DeleteContainer(_ context.Context, req *connect.Request[dkrpc.DeleteContainerRequest]) (*connect.Response[dkrpc.DeleteContainerResponse], error) {
	containerId := req.Msg.GetContainerId()
	log.Debug().Str("Container ID", req.Msg.GetContainerId()).Msgf("Recivied delete container request")

	if containerId == "" {
		return nil, errors.New("no container Id found")
	}

	res := connect.NewResponse(&dkrpc.DeleteContainerResponse{})
	return res, nil
}

func (dk *DockerServer) StopContainer(_ context.Context, req *connect.Request[dkrpc.StopContainerRequest]) (*connect.Response[dkrpc.StopContainerResponse], error) {
	combinedId := req.Msg.GetCombinedId()
	err := dockerclient.HandleStopContainerReq(dk.clientList, combinedId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&dkrpc.StopContainerResponse{})
	return res, nil
}
func (dk *DockerServer) CreateNewImage(_ context.Context, req *connect.Request[dkrpc.NewImageRequest]) (*connect.Response[dkrpc.NewImageResponse], error) {
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
func (dk *DockerServer) ListImages(_ context.Context, _ *connect.Request[dkrpc.ListImageRequest]) (*connect.Response[dkrpc.ListImageResponse], error) {
	images := dockerclient.HandleListImagesReq(dk.clientList)
	res := connect.NewResponse(&dkrpc.ListImageResponse{Images: images})
	return res, nil
}

func (dk *DockerServer) ListContainers(_ context.Context, _ *connect.Request[dkrpc.ListContainersRequest]) (*connect.Response[dkrpc.ListContainersResponse], error) {
	containerList := dockerclient.HandleListContainerReq(dk.clientList)
	res := connect.NewResponse(&dkrpc.ListContainersResponse{Containers: containerList})
	return res, nil
}
