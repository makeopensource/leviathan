package docker

import (
	"fmt"
	"github.com/makeopensource/leviathan/common"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGetLeastJobCountMachineId_100Clients(t *testing.T) {
	manager := &RemoteClientManager{
		Clients: make(map[string]*Machine),
	}

	// Initialize 100 clients
	for i := 0; i < 100; i++ {
		machineID := fmt.Sprintf("machine%d", i)
		manager.Clients[machineID] = &Machine{
			Client:     &DkClient{},
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

func TestRemoteClientManager_DecreaseJobCount(t *testing.T) {
	clients := map[string]*Machine{
		"id1": {ActiveJobs: 1},
	}
	manager := &RemoteClientManager{Clients: clients, mu: sync.RWMutex{}}

	manager.DecreaseJobCount("id1")
	assert.Equal(t, uint64(0), clients["id1"].ActiveJobs)

	manager.DecreaseJobCount("invalid_id")                // Should not panic, but log a warning. We check for no panic.
	assert.Equal(t, uint64(0), clients["id1"].ActiveJobs) // Ensure count isn't changed on bad ID

	// Test decreasing when count is already 0
	manager.DecreaseJobCount("id1")
	assert.Equal(t, uint64(0), clients["id1"].ActiveJobs)
}

func TestRemoteClientManager_GetClientById(t *testing.T) {
	clients := map[string]*Machine{
		"id1": {Client: &DkClient{}}, // Dummy DkClient
	}
	manager := &RemoteClientManager{Clients: clients, mu: sync.RWMutex{}}

	client, err := manager.GetClientById("id1")
	assert.NoError(t, err)
	assert.NotNil(t, client)

	client, err = manager.GetClientById("invalid_id")
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestRemoteClientManager_GetLeastJobCountMachineId(t *testing.T) {
	clients := map[string]*Machine{
		"id1": {ActiveJobs: 2},
		"id2": {ActiveJobs: 0},
		"id3": {ActiveJobs: 1},
	}
	manager := &RemoteClientManager{Clients: clients, mu: sync.RWMutex{}}

	id := manager.GetLeastJobCountMachineId()
	assert.Equal(t, "id2", id)

	// Test case: All machines have the same job count.
	clients = map[string]*Machine{
		"id1": {ActiveJobs: 0},
		"id2": {ActiveJobs: 0},
		"id3": {ActiveJobs: 0},
	}
	manager = &RemoteClientManager{Clients: clients, mu: sync.RWMutex{}}
	id = manager.GetLeastJobCountMachineId()
	assert.NotEmpty(t, id, "Should return an ID even if all counts are equal")
	assert.Contains(t, []string{"id1", "id2", "id3"}, id, "Returned ID should be one of the available IDs")
}

func TestNewSSHClientWithPasswordAuth(t *testing.T) {
	common.InitConfig()
	InitKeyPairFile()

	// when running this test update the config.yaml with the test machine info
	cli := getClientList()
	mName := "test"
	machine, ok := cli[mName]
	if !ok {
		t.Fatalf("machine %s not configured", mName)
		return
	}

	client, err := NewSSHClientWithPasswordAuth(machine)
	if err != nil {
		t.Fatalf("failed create remote docker client %v", err)
		return
	}

	images, err := client.ListImages()
	if err != nil {
		t.Fatalf("failed list images %v", err)
		return
	}

	for _, image := range images {
		t.Log(image)
	}
}
