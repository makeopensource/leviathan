package dockerclient

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	dktypes "github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

func HandleStartContainerReq(clientList []*client.Client, combinedId string) error {
	containerId, machineId, err := parseCombinedID(combinedId)
	if err != nil {
		return err
	}

	err = StartContainer(clientList[machineId], containerId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to stop container at machine: %s with id", containerId)
		return errors.New("Failed to start container")
	}
	return nil
}

func HandleStopContainerReq(clientList []*client.Client, combinedId string) error {
	containerId, machineId, err := parseCombinedID(combinedId)
	if err != nil {
		return err
	}

	err = StopContainer(clientList[machineId], containerId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to stop container at machine: %s with id", containerId)
		return errors.New("Failed to stop container")
	}
	return nil
}

func HandleListImagesReq(clientList []*client.Client) []*dktypes.DockerImage {
	var result []*dktypes.DockerImage

	for ind, cli := range clientList {
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
			Id:       strconv.Itoa(ind),
			Metadata: images,
		})
	}
	return result
}

func HandleNewImageReq(filename string, contents []byte, imageTag string, clientList []*client.Client) error {
	if len(filename) == 0 {
		return errors.New("filename is missing")
	} else if len(contents) == 0 {
		return errors.New("file contents are missing")
	} else if len(imageTag) == 0 {
		return errors.New("imagetag is missing")
	}

	uploadDir := "./uploads"
	fullpath := fmt.Sprintf("%s/%s", uploadDir, filename)

	err := saveDockerfile(fullpath, contents)
	if err != nil {
		return err
	}

	for _, cli := range clientList {
		err := BuildImageFromDockerfile(cli, fullpath, imageTag)
		if err != nil {
			info, err := cli.Info(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("failed to get server info")
				return errors.New("failed to get server info")
			}
			log.Error().Err(err).Msgf("Error building image for %s", info.Name)
			return errors.New("failed to create image for a client")
		}
	}

	return nil
}

func HandleListContainerReq(clientList []*client.Client) []*dktypes.DockerContainer {
	var result []*dktypes.DockerContainer
	for ind, cli := range clientList {
		containers, err := ListContainers(cli, strconv.Itoa(ind))
		if err != nil {
			info, err := cli.Info(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("failed to get docker server info")
				continue
			}
			log.Error().Err(err).Msgf("Error listing containers for %s", info.Name)
			continue
		}

		result = append(result, &dktypes.DockerContainer{
			Id:       strconv.Itoa(ind),
			Metadata: containers,
		})

	}
	return result
}

func HandleCreateContainerReq(clientList []*client.Client, imageTag string, studentCode string) {
}
