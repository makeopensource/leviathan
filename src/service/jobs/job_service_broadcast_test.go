package jobs

import (
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var statusList = []models.JobStatus{models.Complete, models.Failed, models.Canceled}

func TestBroadcastJobs(t *testing.T) {
	SetupTest()

	numJobs := 5
	var jobList []*models.Job

	for i := 0; i < numJobs; i++ {
		jobList = append(jobList, &models.Job{JobId: uuid.New().String()})
	}

	// run multiple listeners for the same job
	for _, job := range jobList {
		t.Run(job.JobId, func(t *testing.T) {
			t.Parallel()
			testSingleBroadcastJob(t, job)
		})
	}

}

func testSingleBroadcastJob(t *testing.T, job *models.Job) {
	dupJobs := 4 // multiple listeners for the same job

	// Wait group for stream listeners
	var wg sync.WaitGroup
	wg.Add(dupJobs)

	var streams []chan *models.Job
	for i := 0; i < dupJobs; i++ {
		streams = append(streams, JobTestService.GetJobInfoChannel(job.JobId))
	}
	defer func() {
		for _, stream := range streams {
			JobTestService.broadcastCh.Unsubscribe(job.JobId, stream)
		}
	}()

	for _, stream := range streams {
		go func() {
			defer wg.Done()
			listenToMessages(stream, t, job)
		}()
	}

	// after 2 secs, update some random value
	time.AfterFunc(time.Second*2, func() {
		log.Info().Msgf("changing job")

		job.StatusMessage = "change after 2"
		res := JobTestService.db.Model(job).Save(job)
		if res.Error != nil {
			t.Fatal(res.Error)
		}
	})

	// after 4 secs, change status to a random final status
	time.AfterFunc(time.Second*4, func() {
		log.Info().Msgf("Job %v finished", job.JobId)

		job.Status = statusList[rand.Intn(3)]
		res := JobTestService.db.Model(job).Save(job)
		if res.Error != nil {
			t.Fatal(res.Error)
		}
	})

	// Wait for listener goroutines to finish
	wg.Wait()
}

func listenToMessages(dupStream chan *models.Job, t *testing.T, job *models.Job) {
	for {
		select {
		case info, ok := <-dupStream:
			if ok {
				assert.Equal(t, info.JobId, job.JobId)
				assert.Equal(t, info.StatusMessage, job.StatusMessage)
			} else {
				// exit loop, channel is done streaming
				return
			}
		}
	}
}
