package service

import (
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
)

func InitServices() (*docker.DkService, *jobs.JobService) {
	dkService := docker.NewDockerServiceWithClients()
	jobService := jobs.NewJobServiceWithDeps(dkService)
	return dkService, jobService
}
