package api

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	dkrpc "github.com/makeopensource/leviathan/generated/docker_rpc/v1"
	"github.com/makeopensource/leviathan/service/dockerclient"
	"github.com/rs/zerolog/log"
)

type DockerServer struct {
	service *dockerclient.DockerService
}

func (dk *DockerServer) CreateContainer(_ context.Context, req *connect.Request[dkrpc.CreateContainerRequest]) (*connect.Response[dkrpc.CreateContainerResponse], error) {
	machineID := "ca068799-ce39-472e-a442-6c30a73cbafe"
	jobId := "freeddrf444"
	imageTag := req.Msg.GetImageTag()

	_, err := dk.service.CreateContainerReq(machineID, jobId, imageTag)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&dkrpc.CreateContainerResponse{})
	return res, nil
}

func (dk *DockerServer) StartContainer(_ context.Context, req *connect.Request[dkrpc.StartContainerRequest]) (*connect.Response[dkrpc.StartContainerResponse], error) {
	err := dk.service.StartContainerReq(req.Msg.GetCombinedId())
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
	err := dk.service.StopContainerReq(combinedId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&dkrpc.StopContainerResponse{})
	return res, nil
}

func (dk *DockerServer) GetContainerLogs(_ context.Context, req *connect.Request[dkrpc.GetContainerLogRequest], responseStream *connect.ServerStream[dkrpc.GetContainerLogResponse]) error {
	err := dk.service.StreamContainerLogs(req.Msg.GetCombinedId(), responseStream)
	if err != nil {
		return err
	}
	return nil
}

func (dk *DockerServer) CreateNewImage(_ context.Context, req *connect.Request[dkrpc.NewImageRequest]) (*connect.Response[dkrpc.NewImageResponse], error) {
	filename := req.Msg.File.GetFilename()
	contents := req.Msg.File.GetContent()
	imageTag := req.Msg.GetImageTag()

	err := dk.service.NewImageReq(filename, contents, imageTag)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&dkrpc.NewImageResponse{})
	return res, nil
}
func (dk *DockerServer) ListImages(_ context.Context, _ *connect.Request[dkrpc.ListImageRequest]) (*connect.Response[dkrpc.ListImageResponse], error) {
	images := dk.service.ListImagesReq()
	res := connect.NewResponse(&dkrpc.ListImageResponse{Images: images})
	return res, nil
}

func (dk *DockerServer) ListContainers(_ context.Context, _ *connect.Request[dkrpc.ListContainersRequest]) (*connect.Response[dkrpc.ListContainersResponse], error) {
	containerList := dk.service.ListContainerReq()
	res := connect.NewResponse(&dkrpc.ListContainersResponse{Containers: containerList})
	return res, nil
}
