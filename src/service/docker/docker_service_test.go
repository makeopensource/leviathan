package docker

import (
	"github.com/docker/docker/api/types/container"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/common"
	"testing"
)

func TestCopyToContainer(t *testing.T) {
	SetupTest()

	machine, err := DkTestService.ClientManager.GetClientById(DkTestService.ClientManager.GetLeastJobCountMachineId())
	if err != nil {
		t.Fatalf("%v", err)
	}

	ifg := uuid.New()

	contId, err := machine.CreateNewContainer(ifg.String(), ImageName, "", container.Resources{})
	if err != nil {
		t.Fatalf("%v", err)
	}

	dir, err := common.CreateTmpJobDir(ifg.String(), map[string][]byte{
		"test.txt": []byte("test test"),
	})
	if err != nil {
		t.Fatalf("%v", err)
		return
	}

	// just copy and check if it succeeds
	err = machine.CopyToContainer(contId, dir)
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = machine.RemoveContainer(contId, true, true)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
