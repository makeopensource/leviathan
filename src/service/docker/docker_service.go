package docker

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	cont "github.com/docker/docker/api/types/container"
	dktypes "github.com/makeopensource/leviathan/generated/docker_rpc/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type DkService struct {
	ClientManager *RemoteClientManager
}

func NewDockerService(clientList *RemoteClientManager) *DkService {
	return &DkService{ClientManager: clientList}
}

func (service *DkService) StartContainerReq(combinedId string) error {
	containerId, machineId, err := ParseCombinedID(combinedId)
	if err != nil {
		return err
	}

	machine, err := service.ClientManager.GetClientById(machineId)
	if err != nil {
		return err
	}

	err = machine.StartContainer(containerId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to stop container at machine: %s with id", containerId)
		return errors.New("Failed to start container")
	}
	return nil
}

func (service *DkService) StopContainerReq(combinedId string) error {
	containerId, machineId, err := ParseCombinedID(combinedId)
	if err != nil {
		return err
	}

	machine, err := service.ClientManager.GetClientById(machineId)
	if err != nil {
		return err
	}

	err = machine.StopContainer(containerId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to stop container at machine: %s with id", containerId)
		return errors.New("Failed to stop container")
	}
	return nil
}

func (service *DkService) ListImagesReq() []*dktypes.DockerImage {
	var result []*dktypes.DockerImage

	for machineID, cli := range service.ClientManager.Clients {
		images, err := cli.Client.ListImages()
		if err != nil {
			info, err := cli.Client.Client.Info(context.Background())
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

func (service *DkService) NewImageReq(filename string, contents []byte, imageTag string) error {
	if len(filename) == 0 {
		return fmt.Errorf("filename is missing")
	} else if len(contents) == 0 {
		return fmt.Errorf("file contents are missing")
	} else if len(imageTag) == 0 {
		return fmt.Errorf("imagetag is missing")
	}

	fullpath, err := SaveDockerfile(filename, contents)
	if err != nil {
		return err
	}

	for _, item := range service.ClientManager.Clients {
		err := item.Client.BuildImageFromDockerfile(fullpath, imageTag)
		if err != nil {
			info, err := item.Client.Client.Info(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("failed to get server info")
				return fmt.Errorf("failed to get server info")
			}
			log.Error().Err(err).Msgf("Error building image for %s", info.Name)
			return fmt.Errorf("failed to create image for a web_gen")
		}
	}

	return nil
}

func (service *DkService) ListContainerReq() []*dktypes.DockerContainer {
	var result []*dktypes.DockerContainer
	for machineID, cli := range service.ClientManager.Clients {
		info, err := cli.Client.Client.Info(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to get docker server info")
			continue
		}
		containers, err := cli.Client.ListContainers(info.ID)
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

func (service *DkService) StreamContainerLogs(combinedId string, responseStream *connect.ServerStream[dktypes.GetContainerLogResponse]) error {
	return nil
}

func (service *DkService) CreateContainerReq(machineId string, jobId string, imageTag string) (string, error) {
	if machineId == "" {
		return "", errors.New("machineId is empty or missing")
	}
	if jobId == "" {
		return "", errors.New("jobID is empty or missing")
	}
	if imageTag == "" {
		return "", errors.New("imageTag is empty or missing")
	}

	machine, err := service.ClientManager.GetClientById(machineId)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to get machine info")
		return "", fmt.Errorf("failed to get machine info")
	}

	resources := cont.Resources{
		Memory:   512 * 1000000,
		NanoCPUs: 2 * 1000000000,
	}

	containerID, err := machine.CreateNewContainer(jobId, imageTag, resources, "")
	if err != nil {
		log.Error().Err(err).Msgf("Failed to create container for job %s", jobId)
		return "", err
	}

	return containerID, nil
}
