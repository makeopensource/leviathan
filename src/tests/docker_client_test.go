package tests

import (
	"fmt"
	"github.com/makeopensource/leviathan/service/docker"
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

	// Test for consistent selection with same job count
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
