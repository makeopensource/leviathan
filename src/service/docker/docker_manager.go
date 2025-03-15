package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
	"github.com/go-viper/mapstructure/v2"
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"sync"
)

type MachineStatus struct {
	Client     *DkClient
	ActiveJobs uint64
}

type RemoteClientManager struct {
	Clients map[string]*MachineStatus
	mu      sync.RWMutex
}

func GetClientList() []models.MachineOptions {
	var allMachines []models.MachineOptions

	// Get all settings
	allSettings := viper.Get("clients.ssh")
	clients, ok := allSettings.(map[string]interface{})
	if !ok {
		log.Warn().Msg("clients.ssh not configured, ssh docker clients will not be used")
		return allMachines
	}

	for name, info := range clients {
		var options models.MachineOptions

		// Decode using mapstructure
		err := mapstructure.Decode(info, &options)
		if err != nil {
			log.Warn().Err(err).Msgf("Error decoding configuration structure for %s", name)
			continue
		}
		// Set the name manually since it's not part of the nested structure
		options.Name = name
		if options.Enable {
			allMachines = append(allMachines, options)
			log.Info().Any("options", options).Msgf("found machine config: %s", name)
		} else {
			log.Debug().Any("options", options).Msgf("found machine config: %s, but it was disabled", name)
		}
	}

	return allMachines
}

func InitDockerClients() *RemoteClientManager {
	// contains clients loaded from config
	untestedClientList := GetClientList()
	// contains final connected list
	clientList := map[string]*MachineStatus{}

	for _, machine := range untestedClientList {
		connStr := fmt.Sprintf("%s@%s:%d", machine.User, machine.Host, machine.Port)
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
			ActiveJobs: 0,
		}
	}

	if common.EnableLocalDocker.GetBool() {
		localClient, err := NewLocalClient()
		if err != nil {
			log.Error().Err(err).Msg("Failed to setup local docker client")
		}

		info, err := testClientConn(localClient.Client)
		if err != nil {
			log.Warn().Err(err).Msgf("Client failed to connect: localdocker")
		} else {
			clientList[info.ID] = &MachineStatus{
				Client:     localClient,
				ActiveJobs: 0,
			}
		}
	} else {
		log.Warn().Msgf("Local docker is disabled in config")
	}

	if len(clientList) == 0 {
		// machines should always be available
		log.Fatal().Msgf("No docker clients connected, check your config")
	}

	return &RemoteClientManager{Clients: clientList, mu: sync.RWMutex{}}
}

// GetLeastJobCountMachineId decides which machine the job run on
// for now least jobs running will be picked
// should be changed to factor in machine resources and load balance accordingly
func (man *RemoteClientManager) GetLeastJobCountMachineId() string {
	man.mu.Lock()

	var minCount = ^uint64(0) // Initialize with the maximum possible value
	machineInd := ""
	for i, v := range man.Clients {
		if v.ActiveJobs < minCount {
			minCount = v.ActiveJobs
			machineInd = i
		}
	}

	// Handle the case where all machines have the same (likely 0) active jobs.
	// In this scenario, the loop might not update machineInd if all values are equal.
	if machineInd == "" {
		// Pick the first machine in the list as a default if none were selected.
		for i := range man.Clients {
			machineInd = i
			break
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

	mac.ActiveJobs++
}

func (man *RemoteClientManager) DecreaseJobCount(id string) {
	man.mu.Lock()
	defer man.mu.Unlock()

	mac, exists := man.Clients[id]
	if !exists {
		log.Warn().Msgf("Invalid machine ID: %s", id)
		return
	}

	if mac.ActiveJobs > 0 {
		mac.ActiveJobs--
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
