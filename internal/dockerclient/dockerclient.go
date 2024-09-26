package dockerclient

import (
	"fmt"
	"net/http"
	"os"

	"github.com/docker/cli/cli/connhelper"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

// NewSSHClient creates a new SSH based client
func NewSSHClient(connectionString string) (*client.Client, error) {
	helper, err := connhelper.GetConnectionHelper(fmt.Sprintf("ssh://%s:22", connectionString))
	if err != nil {
		log.WithFields(log.Fields{"error": err, "connectionString": connectionString}).Error("failed get connectionhelper")
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
		log.WithFields(log.Fields{"error": err, "connectionString": connectionString}).Error("failed create docker client")
		return nil, fmt.Errorf("unable to create docker client")
	}

	return newClient, nil
}

// Build image
func BuildImageFromDockerfile(client *client.Client) error {

	return nil
}

// ListContainer lists all the containers running on host machine
//func ListContainer(c *client.Client) error {
//	containers, err := c.ContainerList(context.Background(), types.ContainerListOptions{All: true})
//	if err != nil {
//		log.WithFields(log.Fields{"error": err}).Error("failed to list containers")
//		return err
//	}
//	// TODO: return containers and print elsewhere
//	if len(containers) > 0 {
//		for _, container := range containers {
//			log.Infof("Container ID: %s", container.ID)
//		}
//	} else {
//		log.Infof("There are no containers running\n")
//	}
//	return nil
//}

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
//
//// RemoveContainer deletes the container of a given ID
//func RemoveContainer(c *client.Client, containerID string, force bool, removeVolumes bool) error {
//	err := c.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: force, RemoveVolumes: removeVolumes})
//	if err != nil {
//		log.WithFields(log.Fields{"error": err, "container_id": containerID}).Error("failed to remove container")
//		return err
//	}
//	return nil
//}
//
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
//
//// StartContainer starts the container of a given ID
//func StartContainer(c *client.Client, containerID string) error {
//	err := c.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
//	if err != nil {
//		log.WithFields(log.Fields{"error": err, "container_id": containerID}).Error("failed to start container")
//		return err
//	}
//	return nil
//}
//
//// StopContainer stops the container of a given ID
//func StopContainer(c *client.Client, containerID string) error {
//	err := c.ContainerStop(context.Background(), containerID, nil)
//	if err != nil {
//		log.WithFields(log.Fields{"error": err, "container_id": containerID}).Error("failed to stop container")
//		return err
//	}
//	return nil
//}
//
//func TailContainerLogs(ctx context.Context, c *client.Client, containerID string) error {
//	reader, err := c.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
//	if err != nil {
//		log.WithFields(log.Fields{"error": err, "container_id": containerID}).Error("failed to get container logs")
//		return err
//	}
//
//	_, err = stdcopy.StdCopy(os.Stdout, os.Stdout, reader)
//	if err != nil && err != io.EOF && err != context.Canceled {
//		log.WithFields(log.Fields{"error": err, "container_id": containerID}).Error("failed to show container logs")
//		return err
//	}
//
//	return nil
//}
//
//// PruneContainers clears all containers that are not running
//func PruneContainers(c *client.Client) error {
//	report, err := c.ContainersPrune(context.Background(), filters.Args{})
//	if err != nil {
//		log.WithFields(log.Fields{"error": err}).Error("failed to show container logs")
//		return err
//	}
//	log.WithFields(log.Fields{"container_ids": report.ContainersDeleted}).Infof("containers pruned")
//	return nil
//}
//
//// CopyToContainer copies a specific file directly into the container
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
