package jobs

import (
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/utils"
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
	utils.InitConfig()
	InitServices()
	BuildImage()
}

func InitServices() {
	// utils for services
	db := utils.InitDB()
	fCache := utils.NewLabFilesCache(db)
	clientList := docker.InitDockerClients()

	DkTestService = docker.NewDockerService(clientList)
	JobTestService = NewJobService(db, fCache, DkTestService) // depends on docker service
}

func BuildImage() {
	bytes, err := utils.ReadFileBytes(DockerFilePath)
	if err != nil {
		log2.Fatal("Unable to read Dockerfile " + DockerFilePath)
	}
	err = DkTestService.NewImageReq(filepath.Base(DockerFilePath), bytes, ImageName)
	if err != nil {
		log2.Fatal("Unable to build Dockerfile " + DockerFilePath)
	}
}
