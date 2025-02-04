package service

import (
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/makeopensource/leviathan/service/labs"
	"github.com/makeopensource/leviathan/service/stats"
	"github.com/makeopensource/leviathan/utils"
)

func InitServices() (*docker.DockerService, *labs.LabService, *jobs.JobService, *stats.StatService) {
	// utils for services
	db := utils.InitDB()
	fCache := utils.NewLabFilesCache(db)
	clientList := docker.InitDockerClients()

	dkService := docker.NewDockerService(clientList)
	labService := labs.NewLabService(db, fCache)
	jobService := jobs.NewJobService(db, fCache, dkService) // depends on docker service
	statsService := stats.NewStatsService(db, fCache)

	return dkService, labService, jobService, statsService
}
