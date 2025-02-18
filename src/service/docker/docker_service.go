package docker

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	cont "github.com/docker/docker/api/types/container"
	"github.com/makeopensource/leviathan/common"
	dktypes "github.com/makeopensource/leviathan/generated/docker_rpc/v1"
	"github.com/rs/zerolog/log"
)

type DkService struct {
	ClientManager *RemoteClientManager
}

func NewDockerService(clientList *RemoteClientManager) *DkService {
	return &DkService{ClientManager: clientList}
}

func (service *DkService) StartContainerReq(combinedId string) error {
	containerId, machineId, err := common.ParseCombinedID(combinedId)
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
		return fmt.Errorf("failed to start container")
	}
	return nil
}

func (service *DkService) StopContainerReq(combinedId string) error {
	containerId, machineId, err := common.ParseCombinedID(combinedId)
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
		return fmt.Errorf("failed to stop container")
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

func (service *DkService) NewImageReq(dockerfilePath string, imageTag string) error {
	for _, item := range service.ClientManager.Clients {
		err := item.Client.BuildImageFromDockerfile(dockerfilePath, imageTag)
		if err != nil {
			info, err := item.Client.Client.Info(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("failed to get server info")
				return fmt.Errorf("failed to get server info")
			}
			log.Error().Err(err).Msgf("Error building image for %s", info.Name)
			return fmt.Errorf("failed to create image for client")
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
		return "", fmt.Errorf("machineId is empty or missing")
	}
	if jobId == "" {
		return "", fmt.Errorf("jobID is empty or missing")
	}
	if imageTag == "" {
		return "", fmt.Errorf("imageTag is empty or missing")
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

	containerID, err := machine.CreateNewContainer(jobId, imageTag, resources)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to create container for job %s", jobId)
		return "", err
	}

	return containerID, nil
}
