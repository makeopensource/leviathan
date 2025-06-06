package docker

import (
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/internal/config"
	"sync"
	"testing"
)

var (
	DkTestService *DkService
	setup         sync.Once
)

const (
	ImageName      = "arithmetic-python"
	DockerFilePath = "../../../example/simple-addition/ex-Dockerfile"
)

func TestConcurrentImageBuilds(t *testing.T) {
	setupTest()
	numTimes := 100
	for i := 0; i < numTimes; i++ {
		t.Run(fmt.Sprintf("image_%d", i), func(t *testing.T) {
			t.Parallel()
			DkTestService.BuildNewImageOnAllClients(DockerFilePath, "test")
		})
	}
}

func TestCopyToContainer(t *testing.T) {
	setupTest()

	machine, err := DkTestService.ClientManager.GetClientById(DkTestService.ClientManager.GetLeastJobCountMachineId())
	if err != nil {
		t.Fatalf("%v", err)
	}

	ifg := uuid.New()

	contId, err := machine.CreateNewContainer(ifg.String(), ImageName, "echo hello", container.Resources{})
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = machine.RemoveContainer(contId, true, true)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func setupTest() {
	setup.Do(func() {
		config.LoadConfig()
		initServices()
	})
}

func initServices() {
	clientList := NewRemoteClientManager()
	DkTestService = NewDockerService(clientList)
}
