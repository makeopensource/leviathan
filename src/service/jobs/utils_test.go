package jobs

import (
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	log2 "log"
	"path/filepath"
)

var (
	DkTestService  *docker.DkService
	JobTestService *JobService
)

const (
	ImageName      = "arithmetic-python"
	DockerFilePath = "../../../example/python/simple-addition/ex-Dockerfile"
)

func SetupTest() {
	common.InitConfig()
	InitServices()
	BuildImage()
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

func BuildImage() {
	bytes, err := common.ReadFileBytes(DockerFilePath)
	if err != nil {
		log2.Fatal("Unable to read Dockerfile " + DockerFilePath)
	}
	err = DkTestService.NewImageReq(filepath.Base(DockerFilePath), bytes, ImageName)
	if err != nil {
		log2.Fatal("Unable to build Dockerfile " + DockerFilePath)
	}
}
