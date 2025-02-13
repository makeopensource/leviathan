package tests

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/utils"
	"io"
	"testing"
)

func TestCopyToContainer(t *testing.T) {
	setupTest()
	//
	//machine, err := dkService.ClientManager.GetClientById(dkService.ClientManager.GetLeastJobCountMachineId())
	//if err != nil {
	//	t.Fatalf("%v", err)
	//}
	//
	//ifg := uuid.New()
	//
	//contId, err := machine.CreateNewContainer(ifg.String(), imageName, container.Resources{}, "")
	//if err != nil {
	//	t.Fatalf("%v", err)
	//}

	//jobBytes, err := utils.TarFile()
	//if err != nil {
	//	t.Fatalf("%v", err)
	//}

	//err = machine.Client.CopyToContainer(
	//	context.Background(),
	//	contId,
	//	"/",
	//	jobBytes,
	//	container.CopyToContainerOptions{},
	//)
	//if err != nil {
	//	t.Fatalf("%v", err)
	//}

	//err = machine.StartContainer(contId)
	//if err != nil {
	//	t.Fatalf("%v", err)
	//}
}

func copyToContainer(machine docker.DkClient, ctx context.Context, contId string, fileContent func(tw io.Writer) error, fileContentSize int64, containerFilePath string, fileMode int64) error {
	buffer, err := utils.TarFile(containerFilePath, fileContent, fileContentSize, fileMode)
	if err != nil {
		return err
	}

	err = machine.Client.CopyToContainer(ctx, contId, "/", buffer, container.CopyToContainerOptions{})
	if err != nil {
		return err
	}

	return nil
}
