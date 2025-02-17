package service

import (
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/makeopensource/leviathan/service/labs"
	"github.com/makeopensource/leviathan/service/stats"
)

func InitServices() (*docker.DkService, *labs.LabService, *jobs.JobService, *stats.StatService) {
	// common for services
	db := common.InitDB()
	bc, ctx := models.NewBroadcastChannel()
	// inject broadcast channel to database
	db.WithContext(ctx)

	fCache := models.NewLabFilesCache(db)
	clientList := docker.InitDockerClients()

	dkService := docker.NewDockerService(clientList)
	labService := labs.NewLabService(db, fCache)
	jobService := jobs.NewJobService(db, fCache, bc, dkService) // depends on docker service
	statsService := stats.NewStatsService(db, fCache)

	return dkService, labService, jobService, statsService
}
