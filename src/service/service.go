package service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
	"github.com/makeopensource/leviathan/service/dockerclient"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/makeopensource/leviathan/service/labs"
	"github.com/makeopensource/leviathan/service/stats"
	"github.com/makeopensource/leviathan/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitServices() (*dockerclient.DockerService, *labs.LabService, *jobs.JobService, *stats.StatService) {
	// database
	db := utils.InitDB()
	fCache := utils.NewLabFilesCache(db)

	clientList := initDockerClients()

	dkService := dockerclient.NewDockerService(clientList)
	labService := labs.NewLabService(db, fCache)
	jobService := jobs.NewJobService(db, fCache)
	statsService := stats.NewStatsService(db, fCache)

	return dkService, labService, jobService, statsService
}

func initDockerClients() map[string]*client.Client {
	// contains clients loaded from config
	untestedClientList := utils.GetClientList()
	// contains final connected list
	clientList := map[string]*client.Client{}

	for _, machine := range untestedClientList {
		connStr := fmt.Sprintf("%s@%s", machine.User, machine.Host)
		remoteClient, err := dockerclient.NewSSHClient(connStr)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to setup remote docker client: %s", machine.Name)
			continue
		}

		info, err := testClientConn(remoteClient)
		if err != nil {
			log.Warn().Err(err).Msgf("Client failed to connect: %s", machine.Name)
			continue
		}

		clientList[info.ID] = remoteClient
	}

	if viper.GetBool("clients.enable_local_docker") {
		localClient, err := dockerclient.NewLocalClient()
		if err != nil {
			log.Error().Err(err).Msg("Failed to setup local docker client")
		}

		info, err := testClientConn(localClient)
		if err != nil {
			log.Warn().Err(err).Msgf("Client failed to connect: localdocker")
		}
		clientList[info.ID] = localClient
	} else {
		log.Warn().Msgf("Local docker is disabled in config")
	}

	if len(clientList) == 0 {
		log.Warn().Msgf("No docker clients connected")
	}

	return clientList
}

func testClientConn(client *client.Client) (system.Info, error) {
	info, err := client.Info(context.Background())
	if err != nil {
		return system.Info{}, err
	}

	log.Info().Str("ID", info.ID).Str("Kernel", info.KernelVersion).Msgf("Connected to %v", info.Name)
	return info, nil
}
