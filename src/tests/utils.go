package tests

import (
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/makeopensource/leviathan/utils"
	log2 "log"
	"path/filepath"
)

var (
	dkService  *docker.DkService
	jobService *jobs.JobService
)

const (
	imageName      = "arithmetic-python"
	dockerFilePath = "../../example/python/simple-addition/ex-Dockerfile"
)

func setupTest() {
	utils.InitConfig()
	initServices()
	buildImage()
}

func initServices() {
	// utils for services
	db := utils.InitDB()
	fCache := utils.NewLabFilesCache(db)
	clientList := docker.InitDockerClients()

	dkService = docker.NewDockerService(clientList)
	jobService = jobs.NewJobService(db, fCache, dkService) // depends on docker service
}

func buildImage() {
	bytes, err := utils.ReadFileBytes(dockerFilePath)
	if err != nil {
		log2.Fatal("Unable to read Dockerfile " + dockerFilePath)
	}
	err = dkService.NewImageReq(filepath.Base(dockerFilePath), bytes, imageName)
	if err != nil {
		log2.Fatal("Unable to build Dockerfile " + dockerFilePath)
	}
}
