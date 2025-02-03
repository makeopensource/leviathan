package jobs

import (
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
)

var (
	messageChannel = make(chan *models.JobMessage)
)

const (
	totalJobs = 5
)

func SetupJobQueue() {
	for i := 1; i < totalJobs; i++ {
		go messageProcessors(i)
	}
}

func messageProcessors(workerId int) {
	for msg := range messageChannel {
		log.Info().Msgf("Worker: %d is now processing job: %s", workerId, msg.JobId)

		// get grader metadata
		// need grader file
		// make file
		// image

		// create container

		// transfer files

		// start container

		// stream logs ??
	}
}

func AddJob(mes *models.JobMessage) {
	messageChannel <- mes
}
