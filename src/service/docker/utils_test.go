package docker

import (
	"github.com/makeopensource/leviathan/utils"
	log2 "log"
	"path/filepath"
)

var (
	DkTestService *DkService
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
	clientList := InitDockerClients()
	DkTestService = NewDockerService(clientList)
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
