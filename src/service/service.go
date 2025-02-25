package service

import (
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
)

func InitServices() (*docker.DkService, *jobs.JobService) {
	// common for services
	db := common.InitDB()
	bc, ctx := models.NewBroadcastChannel()
	// inject broadcast channel to database
	db = db.WithContext(ctx)

	clientList := docker.InitDockerClients()

	dkService := docker.NewDockerService(clientList)
	jobService := jobs.NewJobService(db, bc, dkService) // depends on docker service

	return dkService, jobService
}
