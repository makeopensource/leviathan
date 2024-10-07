package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
	"github.com/makeopensource/leviathan/cmd/api"
	"github.com/makeopensource/leviathan/internal/config"
	"github.com/makeopensource/leviathan/internal/dockerclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config.InitConfig()

	clientList := initDockerClients()

	port := "9221"
	srvAddr := fmt.Sprintf(":%s", port)
	mux := api.SetupPaths(clientList)

	log.Info().Msgf("Started server on %s", srvAddr)
	err := http.ListenAndServe(
		srvAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server on %s", srvAddr)
		return
	}
}

func initDockerClients() map[string]*client.Client {
	// contains clients loaded from config
	untestedClientList := config.GetClientList()
	// contains final connected list
	clientList := make(map[string]*client.Client)

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
		log.Warn().Msgf("Local docker is disabled in config.toml")
	}

	if len(clientList) == 0 {
		log.Fatal().Msgf("No docker clients connected")
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
