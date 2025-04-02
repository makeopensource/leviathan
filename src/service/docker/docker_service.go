package docker

import (
	"github.com/rs/zerolog/log"
	"sync"
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

func (dk *DkService) BuildNewImageOnAllClients(dockerfilePath string, imageTag string) {
	var wg sync.WaitGroup

	for name, item := range dk.ClientManager.Clients {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := item.Client.BuildImageFromDockerfile(dockerfilePath, imageTag)
			if err != nil {
				log.Error().Err(err).Msgf("unable to build image: %s for %s", imageTag, name)
				return
			}
			log.Debug().Msgf("image: %s built successfully for machine: %s", imageTag, name)
		}()
	}

	wg.Wait()
}
