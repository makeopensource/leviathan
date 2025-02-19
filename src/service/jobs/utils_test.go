package jobs

import (
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"sync"
)

var (
	DkTestService  *docker.DkService
	JobTestService *JobService
	setupOnce      sync.Once
)

const (
	ImageName      = "arithmetic-python"
	DockerFilePath = "../../../example/simple-addition/ex-Dockerfile"
)

func SetupTest() {
	setupOnce.Do(func() {
		common.InitConfig()
		InitServices()
	})
}

func InitServices() {
	// common for services
	db := common.InitDB()
	bc, ctx := models.NewBroadcastChannel()
	// inject broadcast channel to database
	db = db.WithContext(ctx)

	fCache := models.NewLabFilesCache(db)
	clientList := docker.InitDockerClients()

	DkTestService = docker.NewDockerService(clientList)
	JobTestService = NewJobService(db, fCache, bc, DkTestService) // depends on docker service
}
