package jobs

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMultipleJobIDs(t *testing.T) {
	// Create a new BroadcastChannel
	bc, _ := NewBroadcastChannel()

	// Job IDs
	jobId1 := "job1"
	jobId2 := "job2"

	// Subscribe to job1 and job2
	ch1 := bc.Subscribe(jobId1)
	ch2 := bc.Subscribe(jobId2)

	// Create a job for each jobId
	job1 := &Job{JobId: jobId1, StatusMessage: "data1"}
	job2 := &Job{JobId: jobId2, StatusMessage: "data2"}

	// Broadcast job1 and job2
	bc.Broadcast(job1)
	bc.Broadcast(job2)

	// Test that job1 is received by subscribers of job1
	go func() {
		select {
		case receivedJob := <-ch1:
			assert.Equal(t, job1.JobId, receivedJob.JobId)
			assert.Equal(t, job1.StatusMessage, receivedJob.StatusMessage)
		case <-time.After(time.Second):
			t.Error("Job1 was not received by the channel")
			return
		}
	}()

	// Test that job2 is received by subscribers of job2
	go func() {
		select {
		case receivedJob := <-ch2:
			assert.Equal(t, job2.JobId, receivedJob.JobId)
			assert.Equal(t, job2.StatusMessage, receivedJob.StatusMessage)
		case <-time.After(time.Second):
			t.Error("Job2 was not received by the channel")
			return
		}
	}()
}

func TestBroadcastChannel_Unsubscribe(t *testing.T) {
	bc, _ := NewBroadcastChannel()
	jobID := "job-123"

	// Subscribe and then unsubscribe
	ch := bc.Subscribe(jobID)
	//ch2 := bc.Subscribe(jobID)

	bc.Unsubscribe(jobID)

	// Try broadcasting a message
	job := &Job{JobId: jobID}
	bc.Broadcast(job)

	// Ensure the channel receives no message
	select {
	case <-ch:
		t.Errorf("expected no message but received one")
	case <-time.After(500 * time.Millisecond):
		// Expected behavior, no message should be received
	}

	// Verify the channel is removed from the list
	_, ok := bc.subscribers.Load(jobID)
	if ok {
		channels := []chan *Job{}
		for _, c := range channels {
			if c == ch {
				t.Errorf("channel was not removed from the list")
			}
		}
	}

	// Verify the entire jobID removed from the map
	bc.Unsubscribe(jobID)
	_, ok = bc.subscribers.Load(jobID)
	if ok {
		t.Errorf("no listeners should be listening for this jobId")
	}
}
