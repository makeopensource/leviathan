package dockerclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/cli/cli/connhelper"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// NewSSHClient creates a new SSH based client
func NewSSHClient(connectionString string) (*client.Client, error) {
	helper, err := connhelper.GetConnectionHelper(fmt.Sprintf("ssh://%s:22", connectionString))
	if err != nil {
		log.Error().Err(err).Msgf("connection string: %s", connectionString)
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: helper.Dialer,
		},
	}

	var clientOpts []client.Opt

	clientOpts = append(clientOpts,
		client.WithHTTPClient(httpClient),
		client.WithHost(helper.Host),
		client.WithDialContext(helper.Dialer),
	)

	version := os.Getenv("DOCKER_API_VERSION")

	if version != "" {
		clientOpts = append(clientOpts, client.WithVersion(version))
	} else {
		clientOpts = append(clientOpts, client.WithAPIVersionNegotiation())
	}

	newClient, err := client.NewClientWithOpts(clientOpts...)
	if err != nil {
		log.Error().Err(err).Msgf("failed create docker client connectionString %s", connectionString)
		return nil, fmt.Errorf("unable to create docker client")
	}

	return newClient, nil
}

// Docker image controls

// BuildImageFromDockerfile Build image
func BuildImageFromDockerfile(client *client.Client, dockerfilePath string, imageName string) error {
	dockerFileName := filepath.Base(dockerfilePath)

	// Create a tar archive of the build context
	buildContextTar, err := archive.TarWithOptions(dockerfilePath, &archive.TarOptions{})
	if err != nil {
		return fmt.Errorf("failed to create build context tar: %v", err)
	}

	defer func(buildContextTar io.ReadCloser) {
		err := buildContextTar.Close()
		if err != nil {
			log.Error().Err(err).Msgf("failed to close build context tar")
		}
	}(buildContextTar)

	// Build the Docker image
	resp, err := client.ImageBuild(
		context.Background(),
		buildContextTar,
		types.ImageBuildOptions{
			Context:    buildContextTar,
			Dockerfile: dockerFileName,
		})

	if err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msgf("failed to close Docker image")
		}
	}(resp.Body)

	// Print the build output
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read build output: %v", err)
	}

	log.Info().Msgf("Docker image '%s' built successfully", imageName)
	return nil
}

// ListImages lists all images on the docker client
func ListImages(client *client.Client) ([]ImageInfo, error) {
	imageInfos, err := client.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		log.Error().Err(err).Msgf("failed to list Docker images")
		return nil, err
	}
	log.Debug().Msgf("Docker images listed: %d", len(imageInfos))

	var imageInfoList []ImageInfo
	for _, item := range imageInfos {
		info := ImageInfo{
			repoTags:  item.RepoTags,
			CreatedAt: item.Created,
			Id:        item.ID,
			Size:      item.Size,
		}
		imageInfoList = append(imageInfoList, info)
	}

	return imageInfoList, nil
}

// ListContainers lists containers
func ListContainers(client *client.Client) ([]ContainerInfo, error) {
	containerInfos, err := client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		log.Error().Err(err).Msgf("failed to list Docker images")
		return nil, err
	}
	log.Debug().Msgf("Docker images listed: %d", len(containerInfos))

	var containerInfoList []ContainerInfo
	for _, item := range containerInfos {
		info := ContainerInfo{
			ID:      item.ID,
			Names:   item.Names,
			Image:   item.Image,
			ImageID: item.ImageID,
			Command: item.Command,
			Created: item.Created,
			Ports:   nil,
			IP:      "",
			Labels:  item.Labels,
			State:   item.State,
			Status:  item.Status,
		}

		containerInfoList = append(containerInfoList, info)
	}

	return containerInfoList, nil
}

// CreateNewContainer creates a new container from given image
//func CreateNewContainer(c *client.Client, image string) (string, error) {
//	config := &container.Config{
//		Image: image,
//		Cmd:   []string{"sh", "-c", "su autolab -c \"autodriver -u 100 -f 104857600 -t 900 -o 104857600 autolab\""},
//	}
//	hostConfig := &container.HostConfig{
//		Resources: container.Resources{
//			Memory:   512 * 1000000,
//			NanoCPUs: 2 * 1000000000,
//		},
//	}
//	networkingConfig := &network.NetworkingConfig{}
//	var platform *imagespecs.Platform = nil
//
//	cont, err := c.ContainerCreate(context.Background(), config, hostConfig, networkingConfig, platform, "")
//	if err != nil {
//		if client.IsErrNotFound(err) {
//			log.WithFields(log.Fields{"error": err, "image": image}).Warn("image not found, attempting to pull from registry")
//			if err := PullImage(c, image); err != nil {
//				return "", err
//			}
//			cont, err = c.ContainerCreate(context.Background(), config, hostConfig, networkingConfig, platform, "")
//			if err == nil {
//				return cont.ID, err
//			}
//		}
//		log.WithFields(log.Fields{"error": err, "image": image}).Error("failed to create container")
//		return "", err
//	}
//
//	return cont.ID, nil
//}

//// PullImage clears all containers that are not running
//func PullImage(c *client.Client, image string) error {
//	log.WithFields(log.Fields{"image": image}).Debug("pulling image")
//	out, err := c.ImagePull(context.Background(), image, types.ImagePullOptions{})
//	if err != nil {
//		log.WithFields(log.Fields{"error": err, "image": image}).Error("failed to pull image")
//		return err
//	}
//	defer out.Close()
//
//	response, err := ioutil.ReadAll(out)
//	if err != nil {
//		return err
//	}
//
//	if log.GetLevel() == log.TraceLevel {
//		util.MultiLineResponseTrace(string(response), "ImagePull Response")
//	}
//
//	return nil
//}

// Container controls

// StartContainer starts the container of a given ID
func StartContainer(client *client.Client, containerID string) error {
	err := client.ContainerStart(context.Background(), containerID, container.StartOptions{})
	if err != nil {
		log.Error().Err(err).Msgf("failed to start Docker container")
		return err
	}
	return nil
}

// StopContainer stops the container of a given ID
func StopContainer(client *client.Client, containerID string) error {
	err := client.ContainerStop(context.Background(), containerID, container.StopOptions{})
	if err != nil {
		log.Error().Err(err).Msgf("failed to stop Docker container")
		return err
	}
	return nil
}

// RemoveContainer deletes the container of a given ID
func RemoveContainer(c *client.Client, containerID string, force bool, removeVolumes bool) error {
	err := c.ContainerRemove(
		context.Background(),
		containerID, container.RemoveOptions{
			Force: force, RemoveVolumes: removeVolumes,
		})

	if err != nil {
		log.Error().Err(err).Msgf("failed to remove Docker container")
		return err
	}

	return nil
}

// TailContainerLogs get logs
func TailContainerLogs(ctx context.Context, client *client.Client, containerID string) error {
	reader, err := client.ContainerLogs(ctx, containerID, container.LogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
	if err != nil {
		log.Error().Err(err).Msgf("failed to tail Docker container")
		return err
	}

	_, err = stdcopy.StdCopy(os.Stdout, os.Stdout, reader)
	if err != nil && err != io.EOF && !errors.Is(err, context.Canceled) {
		log.Error().Err(err).Msgf("failed to tail Docker container")
		return err
	}

	return nil
}

// CopyToContainer copies a specific file directly into the container
//func CopyToContainer(c *client.Client, containerID string, filePath string) error {
//	// TODO FIXME - doesnt validate filePath exists
//	archive, err := archive.Tar(filePath, archive.Gzip)
//	if err != nil {
//		log.WithFields(log.Fields{"error": err, "container_id": containerID, "filePath": filePath}).Error("failed to build archive")
//		return err
//	}
//	defer archive.Close()
//
//	config := types.CopyToContainerOptions{
//		AllowOverwriteDirWithFile: true,
//	}
//	err = c.CopyToContainer(context.Background(), containerID, "/home/autolab/", archive, config)
//	if err != nil {
//		log.WithFields(log.Fields{"error": err, "container_id": containerID, "filePath": filePath}).Error("failed to copy files into container")
//		return err
//	}
//	return nil
//}

// general administrative controls

// PruneContainers clears all containers that are not running
//func PruneContainers(c *client.Client) error {
//	report, err := c.ContainersPrune(context.Background(), filters.Args{})
//	if err != nil {
//		log.WithFields(log.Fields{"error": err}).Error("failed to show container logs")
//		return err
//	}
//	log.WithFields(log.Fields{"container_ids": report.ContainersDeleted}).Infof("containers pruned")
//	return nil
//}
