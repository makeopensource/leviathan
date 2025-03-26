package service

import (
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/makeopensource/leviathan/service/labs"
)

func InitServices() (*docker.DkService, *jobs.JobService, *labs.LabService) {
	db, bc := common.InitDB()
	dkService := docker.NewDockerServiceWithClients()
	labService := labs.NewLabService(db, dkService)
	jobService := jobs.NewJobService(db, bc, dkService, labService)
	return dkService, jobService, labService
}
