package jobs

import (
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/models"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"testing"
	"time"
)

var statusList = []models.JobStatus{models.Complete, models.Failed, models.Canceled}

// TODO
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

			time.AfterFunc(1*time.Second, func() {
				job.StatusMessage = "changing job status"
				job.Status = models.Queued
				JobTestService.db.Model(job).Save(job)
			})

			time.AfterFunc(2*time.Second, func() {
				job.StatusMessage = "changing job status"
				job.Status = models.Running
				JobTestService.db.Model(job).Save(job)
			})

			time.AfterFunc(4*time.Second, func() {
				job.StatusMessage = "changing job"
				// random 'finished' status
				job.Status = statusList[rand.IntN(len(statusList))]
				JobTestService.db.Model(job).Save(job)
			})

			jobCh, _ := JobTestService.SubToJob(job.JobId)

			for {
				select {
				case jobFromCh, ok := <-jobCh:
					if ok {
						assert.Equal(t, job.Status, jobFromCh.Status)
						assert.Equal(t, job.StatusMessage, jobFromCh.StatusMessage)
					} else {
						continue
					}
				case <-time.After(15 * time.Second):
					t.Fatal("timed out waiting for job status to change")
				}
			}
		})
	}
}
