package dockerclient

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	cont "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	dktypes "github.com/makeopensource/leviathan/generated/docker_rpc/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

type DockerService struct {
	clientList map[string]*client.Client
}

func NewDockerService(clientList map[string]*client.Client) *DockerService {
	return &DockerService{clientList: clientList}
}

func (service *DockerService) StartContainerReq(combinedId string) error {
	containerId, machineId, err := parseCombinedID(combinedId)
	if err != nil {
		return err
	}

	err = StartContainer(service.clientList[machineId], containerId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to stop container at machine: %s with id", containerId)
		return errors.New("Failed to start container")
	}
	return nil
}

func (service *DockerService) StopContainerReq(combinedId string) error {
	containerId, machineId, err := parseCombinedID(combinedId)
	if err != nil {
		return err
	}

	err = StopContainer(service.clientList[machineId], containerId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to stop container at machine: %s with id", containerId)
		return errors.New("Failed to stop container")
	}
	return nil
}

func (service *DockerService) ListImagesReq() []*dktypes.DockerImage {
	var result []*dktypes.DockerImage

	for machineID, cli := range service.clientList {
		images, err := ListImages(cli)
		if err != nil {
			info, err := cli.Info(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("failed to get docker server info")
				continue
			}
			log.Error().Err(err).Msgf("Error listing images for %s", info.Name)
			continue
		}

		result = append(result, &dktypes.DockerImage{
			Id:       machineID,
			Metadata: images,
		})
	}
	return result
}

func (service *DockerService) NewImageReq(filename string, contents []byte, imageTag string) error {
	if len(filename) == 0 {
		return errors.New("filename is missing")
	} else if len(contents) == 0 {
		return errors.New("file contents are missing")
	} else if len(imageTag) == 0 {
		return errors.New("imagetag is missing")
	}

	uploadDir := "./appdata/uploads"
	fullpath := fmt.Sprintf("%s/%s", uploadDir, filename)

	err := saveDockerfile(fullpath, contents)
	if err != nil {
		return err
	}

	for _, cli := range service.clientList {
		err := BuildImageFromDockerfile(cli, fullpath, imageTag)
		if err != nil {
			info, err := cli.Info(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("failed to get server info")
				return errors.New("failed to get server info")
			}
			log.Error().Err(err).Msgf("Error building image for %s", info.Name)
			return errors.New("failed to create image for a web_gen")
		}
	}

	return nil
}

func (service *DockerService) ListContainerReq() []*dktypes.DockerContainer {
	var result []*dktypes.DockerContainer
	for machineID, cli := range service.clientList {
		info, err := cli.Info(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to get docker server info")
			continue
		}
		containers, err := ListContainers(cli, info.ID)
		if err != nil {
			log.Error().Err(err).Msgf("Error listing containers for %s", info.Name)
			continue
		}

		result = append(result, &dktypes.DockerContainer{
			Id:       machineID,
			Metadata: containers,
		})

	}
	return result
}

func (service *DockerService) StreamContainerLogs(combinedId string, responseStream *connect.ServerStream[dktypes.GetContainerLogResponse]) error {
	containerId, machineId, err := parseCombinedID(combinedId)
	if err != nil {
		return err
	}
	reader, err := TailContainerLogs(context.Background(), service.clientList[machineId], containerId)
	if err != nil {
		return err
	}

	writer := &logStreamWriter{stream: responseStream}
	_, err = stdcopy.StdCopy(writer, writer, reader)
	if err != nil && err != io.EOF && !errors.Is(err, context.Canceled) {
		log.Error().Err(err).Msgf("failed to tail Docker container")
		return errors.New("failed to tail Docker container")
	}

	return nil
}

func (service *DockerService) CreateContainerReq(machineId string, jobId string, imageTag string) (string, error) {
	if machineId == "" {
		return "", errors.New("machineId is empty or missing")
	}
	if jobId == "" {
		return "", errors.New("jobID is empty or missing")
	}
	if imageTag == "" {
		return "", errors.New("imageTag is empty or missing")
	}

	cli := service.clientList[machineId]

	resources := cont.Resources{
		Memory:   512 * 1000000,
		NanoCPUs: 2 * 1000000000,
	}

	containerID, err := CreateNewContainer(cli, jobId, imageTag, resources)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to create container for job %s", jobId)
		return "", err
	}

	return containerID, nil
}
