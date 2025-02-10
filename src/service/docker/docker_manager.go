package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
	"github.com/makeopensource/leviathan/utils"
	"github.com/rs/zerolog/log"
	"sync"
)

type MachineStatus struct {
	Client     *DkClient
	activeJobs uint64
}

type RemoteClientManager struct {
	Clients map[string]*MachineStatus
	mu      sync.RWMutex
}

func InitDockerClients() *RemoteClientManager {
	// contains clients loaded from config
	untestedClientList := utils.GetClientList()
	// contains final connected list
	clientList := map[string]*MachineStatus{}

	for _, machine := range untestedClientList {
		connStr := fmt.Sprintf("%s@%s", machine.User, machine.Host)
		remoteClient, err := NewSSHClient(connStr)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to setup remote docker client: %s", machine.Name)
			continue
		}

		info, err := testClientConn(remoteClient.Client)
		if err != nil {
			log.Warn().Err(err).Msgf("Client failed to connect: %s", machine.Name)
			continue
		}

		clientList[info.ID] = &MachineStatus{
			Client:     remoteClient,
			activeJobs: 0,
		}
	}

	if utils.EnableLocalDocker.GetBool() {
		localClient, err := NewLocalClient()
		if err != nil {
			log.Error().Err(err).Msg("Failed to setup local docker client")
		}

		info, err := testClientConn(localClient.Client)
		if err != nil {
			log.Warn().Err(err).Msgf("Client failed to connect: localdocker")
		}
		clientList[info.ID] = &MachineStatus{
			Client:     localClient,
			activeJobs: 0,
		}
	} else {
		log.Warn().Msgf("Local docker is disabled in config")
	}

	if len(clientList) == 0 {
		// machines should always be available
		log.Fatal().Msgf("No docker clients connected")
	}

	return &RemoteClientManager{Clients: clientList, mu: sync.RWMutex{}}
}

// GetLeastJobCountMachineId decides which machine the job run on
// for now least jobs running will be picked
// should be changed to factor in machine resources and load balance accordingly
func (man *RemoteClientManager) GetLeastJobCountMachineId() string {
	man.mu.Lock()

	var minCount uint64 = 0
	machineInd := ""
	for i, v := range man.Clients {
		if v.activeJobs <= minCount {
			minCount = v.activeJobs
			machineInd = i
		}
	}
	man.mu.Unlock()
	// always unlock before calling this, or it will deadlock
	man.increaseJobCount(machineInd)

	return machineInd
}

func (man *RemoteClientManager) GetClientById(id string) (*DkClient, error) {
	status, exists := man.Clients[id]
	if !exists {
		return nil, fmt.Errorf("invalid machine id: %s", id)
	}

	return status.Client, nil
}

func (man *RemoteClientManager) increaseJobCount(id string) {
	man.mu.Lock()
	defer man.mu.Unlock()

	mac, exists := man.Clients[id]
	if !exists {
		log.Warn().Msgf("Invalid machine ID: %s", id)
		return
	}

	mac.activeJobs++
}

func (man *RemoteClientManager) decreaseJobCount(id string) {
	man.mu.Lock()
	defer man.mu.Unlock()

	mac, exists := man.Clients[id]
	if !exists {
		log.Warn().Msgf("Invalid machine ID: %s", id)
		return
	}

	if mac.activeJobs > 0 {
		mac.activeJobs--
	}
}

func testClientConn(client *client.Client) (system.Info, error) {
	info, err := client.Info(context.Background())
	if err != nil {
		return system.Info{}, err
	}

	log.Info().Str("ID", info.ID).Str("Kernel", info.KernelVersion).Msgf("Connected to %v", info.Name)
	return info, nil
}
