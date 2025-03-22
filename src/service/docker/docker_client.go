package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/makeopensource/leviathan/common"
	dktypes "github.com/makeopensource/leviathan/generated/docker_rpc/v1"
	"github.com/makeopensource/leviathan/models"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/rs/zerolog/log"
	"io"
	"path/filepath"
	"time"
)

type ImageMap = models.Map[string, *models.CountingMutex]

// DkClient a wrapper for the docker client struct, that exposes the commands leviathan needs
type DkClient struct {
	Client *client.Client
	imgMap *ImageMap
}

func NewDkClient(client *client.Client) *DkClient {
	cli := &DkClient{
		Client: client,
		imgMap: &ImageMap{},
	}
	go cleanupImageTagLocks(cli.imgMap)
	return cli
}

// BuildImageFromDockerfile Build image
func (c *DkClient) BuildImageFromDockerfile(dockerfilePath string, tagName string) error {
	// prevent concurrent duplicate image builds
	tagLock := c.imgMap.LoadOrStore(tagName, models.NewCountMutex())
	tagLock.Lock()
	defer tagLock.Unlock()

	dockerfileTar, dockerfile, err := common.TarFile(dockerfilePath)
	if err != nil {
		return fmt.Errorf("failed to tar file %s", dockerfilePath)
	}
	// Build the Docker image
	resp, err := c.Client.ImageBuild(
		context.Background(),
		dockerfileTar,
		types.ImageBuildOptions{
			Context:     dockerfileTar,
			Dockerfile:  dockerfile,
			Tags:        []string{tagName},
			ForceRemove: true, // Removes intermediate containers
			Remove:      true, // Removes intermediate images
		})
	if err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)
	}
	// dispose response
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msgf("failed to close docker image build response")
		}
	}(resp.Body)

	logWriter := &common.LogWriter{LoggerFunc: func(s string) {
		log.Debug().Str("image", tagName).Msgf("%s", s)
	}}
	// Print the build output
	_, err = io.Copy(logWriter, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read build output: %v", err)
	}

	log.Info().Msgf("docker image '%s' built successfully", tagName)
	return nil
}

// ListImages lists all images on the docker web_gen
func (c *DkClient) ListImages() ([]*dktypes.ImageMetaData, error) {
	imageInfos, err := c.Client.ImageList(context.Background(), image.ListOptions{All: true})
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
func (c *DkClient) ListContainers(machineId string) ([]*dktypes.ContainerMetaData, error) {
	containerInfos, err := c.Client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		log.Error().Err(err).Msgf("failed to list Docker images")
		return nil, err
	}

	log.Debug().Msgf("Docker images listed: %d", len(containerInfos))

	var containerInfoList []*dktypes.ContainerMetaData
	return containerInfoList, nil
}

// CreateNewContainer creates a new container from given image
func (c *DkClient) CreateNewContainer(jobUuid, image, jobFolder, entryCmd string, machineLimits container.Resources) (string, error) {
	baseCmd := fmt.Sprintf("cd /home/%s", jobFolder)
	runCommand := fmt.Sprintf("%s && %s", baseCmd, entryCmd)

	config := &container.Config{
		Image: image,
		Labels: map[string]string{
			"con": jobUuid,
		},
		Cmd: []string{"sh", "-c", runCommand},
	}

	hostConfig := &container.HostConfig{
		Resources:  machineLimits,
		AutoRemove: false,
		Binds: []string{
			"/etc/localtime:/etc/localtime:ro", // add a time mount to fix clock skew issue in make
		},
	}
	networkingConfig := &network.NetworkingConfig{}

	var platform *v1.Platform = nil

	cont, err := c.Client.ContainerCreate(
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
func (c *DkClient) StartContainer(containerID string) error {
	err := c.Client.ContainerStart(context.Background(), containerID, container.StartOptions{})
	if err != nil {
		log.Error().Err(err).Msgf("failed to start Docker container")
		return err
	}
	return nil
}

// StopContainer stops the container of a given ID
func (c *DkClient) StopContainer(containerID string) error {
	err := c.Client.ContainerStop(context.Background(), containerID, container.StopOptions{})
	if err != nil {
		log.Error().Err(err).Msgf("failed to stop Docker container")
		return err
	}
	return nil
}

// RemoveContainer deletes the container of a given ID
func (c *DkClient) RemoveContainer(containerID string, force bool, removeVolumes bool) error {
	err := c.Client.ContainerRemove(
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

// CopyToContainer copies a specific dir directly into the container
// stolen from https://github.com/testcontainers/testcontainers-go/blob/f09b3af2cb985a17bd2b2eaaa5d384882ded8e28/docker.go#L633
func (c *DkClient) CopyToContainer(containerID string, submissionDirPath string) error {
	//log.Debug().Msgf("Copying files to container %s", containerDirectory)

	jobBytes, err := common.TarDir(submissionDirPath, 775)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert files to tar")
		return fmt.Errorf("failed to convert files to tar")
	}

	// create the directory under its parent
	parent := filepath.Dir("/home/")

	err = c.Client.CopyToContainer(context.Background(), containerID, parent, jobBytes, container.CopyToContainerOptions{})
	if err != nil {
		log.Error().Err(err).Msgf("failed to copy to container")
		return fmt.Errorf("failed to copy submission to container")
	}

	return nil
}

func (c *DkClient) TailContainerLogs(ctx context.Context, containerID string) (io.ReadCloser, error) {
	reader, err := c.Client.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
		Details:    false,
	})
	if err != nil {
		log.Error().Err(err).Msgf("failed to tail Docker container")
		return nil, err
	}
	return reader, nil
}

// general administrative controls

// PruneContainers clears all containers that are not running
func (c *DkClient) PruneContainers() error {
	report, err := c.Client.ContainersPrune(context.Background(), filters.Args{})
	if err != nil {
		log.Error().Err(err).Msgf("failed to prune Docker container")
		return err
	}

	log.Debug().Msgf("Docker containers pruned: %d", len(report.ContainersDeleted))
	return nil
}

func (c *DkClient) GetContainerStatus(ctx context.Context, contId string) (*container.InspectResponse, error) {
	inspect, err := c.Client.ContainerInspect(ctx, contId)
	if err != nil {
		return nil, err
	}

	return &inspect, nil
}

// cleanup function to run indefinitely, removes locks which have 0 routines waiting on them
func cleanupImageTagLocks(cli *ImageMap) {
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()
	for {
		<-ticker.C
		cli.Range(func(key string, value *models.CountingMutex) bool {
			checkImageLockStatus(key, value, cli)
			// always return true to continue iterating
			return true
		})
	}
}

func checkImageLockStatus(tagName string, mut *models.CountingMutex, imageMap *ImageMap) {
	c := mut.WaitingCount()
	if c == 0 {
		log.Debug().Msgf("removing unused tag: %s from image queue", tagName)
		imageMap.Delete(tagName)
	} else {
		log.Debug().Msgf("tag: %s has locks waiting: %d from image queue", tagName, c)
	}
}
