package cmd

import (
	"github.com/makeopensource/leviathan/internal/config"
	"github.com/makeopensource/leviathan/internal/database"
	"github.com/makeopensource/leviathan/internal/docker"
	fu "github.com/makeopensource/leviathan/internal/file_manager"
	"github.com/makeopensource/leviathan/internal/jobs"
	"github.com/makeopensource/leviathan/internal/labs"
	"github.com/makeopensource/leviathan/pkg/logger"
	"github.com/rs/zerolog/log"
)

func Setup() {
	log.Logger = logger.ConsoleLogger() // logs here are not saved to the log file
	config.LoadConfig()
	// once the log dir and level is set by config,
	// we start a file logger along with the console logger
	log.Logger = logger.FileConsoleLogger(config.LogDir.GetStr(), config.LogLevel.GetStr())
}

func InitServices() (*docker.DkService, *jobs.JobService, *labs.LabService) {
	db, bc := database.NewDatabaseWithGorm()

	dkService := docker.NewDockerServiceWithClients()
	fileManService := fu.NewFileManagerService()
	labService := labs.NewLabService(db, dkService, fileManService)
	jobService := jobs.NewJobService(db, bc, dkService, labService, fileManService)

	return dkService, jobService, labService
}
