package docker

import (
	"context"
	"fmt"
	"github.com/docker/cli/cli/connhelper"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
	"github.com/go-viper/mapstructure/v2"
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh"
	"net/http"
	"sync"
)

type Machine struct {
	Client     *DkClient
	ActiveJobs uint64
}

type MachineMap = map[string]*Machine

type RemoteClientManager struct {
	Clients MachineMap
	mu      sync.RWMutex
}

func NewRemoteClientManager() *RemoteClientManager {
	untestedClientList := getClientList()
	clientList := MachineMap{} // contains final connected list

	for _, machine := range untestedClientList {
		var remoteClient *DkClient
		var err error

		if machine.Password != "" {
			remoteClient, err = NewSSHClientWithPasswordAuth(machine)
		} else {
			remoteClient, err = NewSSHClient(machine)
		}
		if err != nil {
			log.Error().Err(err).Msgf("Failed to setup remote docker client: %s", machine.Name)
			continue
		}

		info, err := testClientConn(remoteClient.Client)
		if err != nil {
			log.Warn().Err(err).Msgf("Remote docker client failed to connect: %s", machine.Name)
			continue
		}

		clientList[info.ID] = &Machine{
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
			clientList[info.ID] = &Machine{
				Client:     localClient,
				ActiveJobs: 0,
			}
		}
	} else {
		log.Info().Msgf("Local docker is disabled in config")
	}

	if len(clientList) < 1 {
		// at least a single machine should always be available
		log.Fatal().Msgf("No docker clients connected, check your config")
	}

	return &RemoteClientManager{Clients: clientList, mu: sync.RWMutex{}}
}

// NewSSHClient connects to a remote docker host using a public private auth.
//
// It is assumed host running leviathan has done the necessary SSH setup, via sshd or other ssh programs
// and the remote host loaded is accessible via SSH
func NewSSHClient(machine models.MachineOptions) (*DkClient, error) {
	connectionString := fmt.Sprintf("%s@%s:%d", machine.User, machine.Host, machine.Port)
	helper, err := connhelper.GetConnectionHelper(fmt.Sprintf("ssh://%s", connectionString))
	if err != nil {
		log.Error().Err(err).Msgf("connection string: %s", connectionString)
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: helper.Dialer,
		},
	}

	newClient, err := client.NewClientWithOpts(
		client.WithHTTPClient(httpClient),
		client.WithHost(helper.Host),
		client.WithDialContext(helper.Dialer),
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed create remote docker client with connectionString %s", connectionString)
		return nil, fmt.Errorf("unable to connect to docker client")
	}

	return NewDkClient(newClient), nil
}

// NewSSHClientWithPasswordAuth connects to a remote docker host using a password.
//
// It is assumed machine models.MachineOptions has the correct password set.
//
// if the public key is empty then leviathan will save the public key on connection
// else it will verify the set public key
func NewSSHClientWithPasswordAuth(machine models.MachineOptions) (*DkClient, error) {
	sshHost := fmt.Sprintf("%s:%d", machine.Host, machine.Port)
	// Create an SSH client configuration.
	config := &ssh.ClientConfig{
		User: machine.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(machine.Password),
		},
		HostKeyCallback: saveHostKey(machine),
	}

	if machine.Publickey != "" {
		log.Debug().Msgf("Verifying public key for %s", machine.Name)
		pubkey, err := stringToPublicKey(machine.Publickey)
		if err != nil {
			return nil, err
		}

		config.HostKeyCallback = ssh.FixedHostKey(pubkey)
	}

	sshClient, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create ssh client: %v", err)
	}

	// Create a Docker client using the custom dialer.
	newClient, err := client.NewClientWithOpts(
		client.WithHost("tcp://"+sshHost), //use the remote host.
		client.WithDialContext(sshDialer(sshClient)),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to docker client")
	}

	return NewDkClient(newClient), nil
}

// NewLocalClient connects to the local docker host.
//
// It is assumed the docker daemon is running and is accessible by leviathan
func NewLocalClient() (*DkClient, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		log.Error().Err(err).Msgf("failed create local docker client")
		return nil, fmt.Errorf("unable to create docker client")
	}

	return NewDkClient(cli), nil
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

// getClientList loads clients from config, if client has 'enable: false' it will be skipped
func getClientList() map[string]models.MachineOptions {
	var machineMap = map[string]models.MachineOptions{}

	// Get all settings
	allSettings := common.ClientsSSH.GetAny()
	clients, ok := allSettings.(map[string]interface{})
	if !ok {
		log.Warn().Msg("clients.ssh not configured, ssh docker clients will not be used")
		return machineMap
	}

	for name, info := range clients {
		var options models.MachineOptions
		err := mapstructure.Decode(info, &options)
		if err != nil {
			log.Warn().Err(err).Msgf("Error decoding configuration structure for %s", name)
			continue
		}

		options.Name = name // Set the name manually since it's not part of the nested structure
		if options.Enable {
			machineMap[name] = options
			log.Info().Any("options", options).Msgf("found machine config: %s", name)
		} else {
			log.Debug().Any("options", options).Msgf("found machine config: %s, but it was disabled", name)
		}
	}

	return machineMap
}

func testClientConn(client *client.Client) (system.Info, error) {
	info, err := client.Info(context.Background())
	if err != nil {
		return system.Info{}, err
	}

	log.Info().Str("ID", info.ID).Str("Kernel", info.KernelVersion).Msgf("Connected to %v", info.Name)
	return info, nil
}
