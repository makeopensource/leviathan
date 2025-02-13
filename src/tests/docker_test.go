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

func TestGetLeastJobCountMachineId_100Clients(t *testing.T) {
	manager := &docker.RemoteClientManager{
		Clients: make(map[string]*docker.MachineStatus),
	}

	// Initialize 100 clients
	for i := 0; i < 100; i++ {
		machineID := fmt.Sprintf("machine%d", i)
		manager.Clients[machineID] = &docker.MachineStatus{
			Client:     &docker.DkClient{},
			ActiveJobs: uint64(i % 5), // Distribute jobs a bit (0-4)
		}
	}

	for i := 0; i < 1000; i++ { // Run multiple times to simulate real-world usage
		machineID := manager.GetLeastJobCountMachineId()

		if _, ok := manager.Clients[machineID]; !ok {
			t.Errorf("Returned machine ID %s does not exist in Clients map", machineID)
		}
	}

	// TODO Test for consistent selection with same job count
	//for k := range manager.Clients {
	//	manager.Clients[k].ActiveJobs = 0 // reset all to 0
	//}
	//firstCall := manager.GetLeastJobCountMachineId()
	//for i := 0; i < 10; i++ {
	//	if manager.GetLeastJobCountMachineId() != firstCall {
	//		t.Errorf("Inconsistent machine selection with same job count")
	//	}
	//}
}

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
