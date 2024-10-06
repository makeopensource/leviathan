package dockerclient

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	dktypes "github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strconv"
)

func HandleCreateContainerReq(clientList []*client.Client, imageTag string, studentCode string) {

}

func HandleStartContainerReq(clientList []*client.Client, containerId string) {

}

func HandleStopContainerReq(clientList []*client.Client, containerId string) {

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

func saveDockerfile(fullPath string, contents []byte) error {

	log.Debug().Str("filename", fullPath).Msgf("Recivied new container request")

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		log.Error().Err(err).Msgf("Failed to create file and folder at %s", fullPath)
		return err
	}

	err := os.WriteFile(fullPath, contents, 0644)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to write contents to file")
		return err
	}

	return nil
}

//func HandleListContainerReq(clientList []*client.Client) [][]ContainerInfo {
//	var result [][]ContainerInfo
//	for _, cli := range clientList {
//		containers, err := ListContainers(cli)
//		if err != nil {
//			info, err := cli.Info(context.Background())
//			if err != nil {
//				log.Error().Err(err).Msg("failed to get docker server info")
//				continue
//			}
//			log.Error().Err(err).Msgf("Error listing containers for %s", info.Name)
//			continue
//		}
//
//		result = append(result, containers)
//	}
//	return result
//}
