package tests

//
//import (
//	"context"
//	"github.com/docker/docker/api/types"
//	"github.com/docker/docker/api/types/container"
//	"github.com/docker/docker/pkg/archive"
//	"github.com/makeopensource/leviathan/service/docker"
//	"github.com/makeopensource/leviathan/utils"
//	"github.com/rs/zerolog/log"
//	"io"
//	"testing"
//)
//
//func TestCopyContainer(t *testing.T) {
//	utils.InitConfig()
//
//	client, err := docker.NewLocalClient()
//	if err != nil {
//		t.Fatal(err)
//		return
//	}
//
//	newContainer, err := client.CreateNewContainer("test", "arithmetic-python", container.Resources{})
//	if err != nil {
//		t.Fatalf("Failed to copy to container: %v", err)
//	}
//
//	const containerDirectory = "/home/autolab/"
//	const filePath = "appdata/submissions/grader.py"
//
//	log.Debug().Msgf("Copying file %s to container %s", filePath, containerDirectory)
//
//	archiveData, err := archive.Tar(filePath, archive.Gzip)
//	if err != nil {
//		log.Error().Err(err).Msgf("failed to archive %s", filePath)
//		t.Fatalf(err.Error())
//	}
//	defer func(archive io.ReadCloser) {
//		err := archive.Close()
//		if err != nil {
//			log.Error().Err(err).Msgf("failed to close archive")
//		}
//	}(archiveData)
//
//	srcDir := "path/to/your/directory"
//	tarReader, err := createTarArchive(srcDir) // Implement createTarArchive function
//	if err != nil {
//		panic(err)
//	}
//	defer tarReader.Close()
//
//	err = client.Client.CopyToContainer(context.Background(), newContainer, containerDirectory, tarReader, types.CopyToContainerOptions{
//		AllowOverwriteDirWithFile: true,
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	if err != nil {
//		t.Fatalf("Failed to copy to container: %v", err)
//	}
//
//	err = client.Client.ContainerStart(context.Background(), newContainer, container.StartOptions{})
//	if err != nil {
//		t.Fatalf("Failed to start container: %v", err)
//	}
//
//}
