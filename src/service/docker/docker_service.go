package docker

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
)

type DkService struct {
	ClientManager *RemoteClientManager
}

func NewDockerService(clientList *RemoteClientManager) *DkService {
	return &DkService{ClientManager: clientList}
}

func NewDockerServiceWithClients() *DkService {
	return &DkService{ClientManager: NewRemoteClientManager()}
}

func (service *DkService) BuildNewImageOnAllClients(dockerfilePath string, imageTag string) error {
	for _, item := range service.ClientManager.Clients {
		err := item.Client.BuildImageFromDockerfile(dockerfilePath, imageTag)
		if err != nil {
			info, err := item.Client.Client.Info(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("failed to get server info")
				return fmt.Errorf("failed to get server info")
			}
			log.Error().Err(err).Msgf("Error building image for %s", info.Name)
			return fmt.Errorf("failed to create image for client")
		}
	}

	return nil
}
