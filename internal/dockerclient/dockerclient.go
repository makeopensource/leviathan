package dockerclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/cli/cli/connhelper"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/stdcopy"
	dktypes "github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1"
	"github.com/makeopensource/leviathan/internal/util"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
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
		client.WithAPIVersionNegotiation(),
	)

	newClient, err := client.NewClientWithOpts(clientOpts...)
	if err != nil {
		log.Error().Err(err).Msgf("failed create remote docker client with connectionString %s", connectionString)
		return nil, fmt.Errorf("unable to connect to docker client")
	}

	return newClient, nil
}

// NewLocalClient create a new client based locally
func NewLocalClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Error().Err(err).Msgf("failed create local docker client")
		return nil, fmt.Errorf("unable to create docker client")
	}

	return cli, nil
}

// Docker image controls

// BuildImageFromDockerfile Build image
func BuildImageFromDockerfile(client *client.Client, dockerfilePath string, tagName string) error {
	_, err := os.Stat(dockerfilePath)
	if err != nil {
		log.Error().Err(err).Msgf("failed to stat path %s", dockerfilePath)
		return err
	}

	dockerfileTar, dockerfile := ConvertToTar(dockerfilePath)
	// Build the Docker image
	resp, err := client.ImageBuild(
		context.Background(),
		dockerfileTar,
		types.ImageBuildOptions{
			Context:    dockerfileTar,
			Dockerfile: dockerfile,
			Tags:       []string{tagName},
		})
	if err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)
	}
	// dispose response
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

	log.Info().Msgf("Docker image '%s' built successfully", tagName)
	return nil
}

// ListImages lists all images on the docker client
func ListImages(client *client.Client) ([]*dktypes.ImageMetaData, error) {
	imageInfos, err := client.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		log.Error().Err(err).Msgf("failed to list Docker images")
		return nil, err
	}
	log.Debug().Msgf("Docker images listed: %d", len(imageInfos))

	var imageInfoList []*dktypes.ImageMetaData
	for _, item := range imageInfos {
		info := dktypes.ImageMetaData{
			RepoTags:  item.RepoTags,
			CreatedAt: item.Created,
			Id:        item.ID,
			Size:      item.Size,
		}
		imageInfoList = append(imageInfoList, &info)
	}
	return imageInfoList, nil
}

// ListContainers lists containers
func ListContainers(client *client.Client, machineId string) ([]*dktypes.ContainerMetaData, error) {
	containerInfos, err := client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		log.Error().Err(err).Msgf("failed to list Docker images")
		return nil, err
	}
	log.Debug().Msgf("Docker images listed: %d", len(containerInfos))

	var containerInfoList []*dktypes.ContainerMetaData
	for _, item := range containerInfos {
		info := dktypes.ContainerMetaData{
			Id:             util.EncodeID(machineId, item.ID),
			ContainerNames: item.Names,
			Image:          item.Image,
			Status:         item.Status,
			State:          item.State,
		}

		containerInfoList = append(containerInfoList, &info)
	}

	return containerInfoList, nil
}

// CreateNewContainer creates a new container from given image
func CreateNewContainer(client *client.Client, jobUuid string, image string, entryPointCmd []string, machineLimits container.Resources) (string, error) {
	config := &container.Config{
		Image: image,
		Labels: map[string]string{
			"con": jobUuid,
		},
		Cmd: entryPointCmd,
	}
	hostConfig := &container.HostConfig{
		Resources:  machineLimits,
		AutoRemove: true,
	}
	networkingConfig := &network.NetworkingConfig{}

	var platform *v1.Platform = nil

	cont, err := client.ContainerCreate(
		context.Background(),
		config,
		hostConfig,
		networkingConfig,
		platform,
		jobUuid,
	)

	if err != nil {
		// maybe pull image if it errors
		log.Error().Err(err).Str("image", image).Msgf("failed to create Docker container")
		return "", err
	}

	return cont.ID, nil
}

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

// I/O within containers

// CopyToContainer copies a specific file directly into the container
func CopyToContainer(client *client.Client, containerID string, filePath string) error {
	const containerDirectory = "/home/autolab/"

	log.Debug().Msgf("Copying file %s to container %s", filePath, containerDirectory)

	_, err := os.Stat(filePath)
	if err != nil {
		log.Error().Err(err).Msgf("failed to stat path %s", filePath)
		return err
	}

	archiveData, err := archive.Tar(filePath, archive.Gzip)
	if err != nil {
		log.Error().Err(err).Msgf("failed to archive %s", filePath)
		return err
	}
	defer func(archive io.ReadCloser) {
		err := archive.Close()
		if err != nil {
			log.Error().Err(err).Msgf("failed to close archive")
		}
	}(archiveData)

	config := container.CopyToContainerOptions{AllowOverwriteDirWithFile: true}
	err = client.CopyToContainer(
		context.Background(),
		containerID,
		containerDirectory,
		archiveData,
		config,
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed to copy to Docker container")
		return err
	}
	return nil
}

// TailContainerLogs get logs TODO
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

// general administrative controls

// PruneContainers clears all containers that are not running
func PruneContainers(c *client.Client) error {
	report, err := c.ContainersPrune(context.Background(), filters.Args{})
	if err != nil {
		log.Error().Err(err).Msgf("failed to prune Docker container")
		return err
	}

	log.Debug().Msgf("Docker containers pruned: %d", len(report.ContainersDeleted))
	return nil
}
