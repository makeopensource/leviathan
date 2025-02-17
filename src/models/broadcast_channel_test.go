package models

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestMultipleJobIDs(t *testing.T) {
	// Create a new BroadcastChannel
	bc := NewBroadcastChannel()

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

func TestBroadcastChannel_MultipleListeners(t *testing.T) {
	bc := NewBroadcastChannel()
	jobID := "job-123"

	// Create multiple listeners
	listenerCount := 3
	listeners := make([]chan *Job, listenerCount)
	for i := 0; i < listenerCount; i++ {
		listeners[i] = bc.Subscribe(jobID)
	}

	// Create a job and broadcast it
	job := &Job{JobId: jobID}
	bc.Broadcast(job)

	// Verify all listeners receive the message
	var wg sync.WaitGroup
	wg.Add(listenerCount)
	for _, ch := range listeners {
		go func(c chan *Job) {
			select {
			case receivedJob := <-c:
				if receivedJob.JobId != jobID {
					t.Errorf("expected job ID %s but got %s", jobID, receivedJob.JobId)
				}
			case <-time.After(1 * time.Second):
				t.Errorf("timed out waiting for message")
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
}

func TestBroadcastChannel_Unsubscribe(t *testing.T) {
	bc := NewBroadcastChannel()
	jobID := "job-123"

	// Subscribe and then unsubscribe
	ch := bc.Subscribe(jobID)
	ch2 := bc.Subscribe(jobID)

	bc.Unsubscribe(jobID, ch)

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
	val, ok := bc.subscribers.Load(jobID)
	if ok {
		channels := val.([]chan *Job)
		for _, c := range channels {
			if c == ch {
				t.Errorf("channel was not removed from the list")
			}
		}
	}

	// Verify the entire jobID removed from the map
	bc.Unsubscribe(jobID, ch2)
	val, ok = bc.subscribers.Load(jobID)
	if ok {
		t.Errorf("no listeners should be listening for this jobId")
	}
}
