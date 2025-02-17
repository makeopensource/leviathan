package models

import (
	"github.com/rs/zerolog/log"
	"sync"
)

type BroadcastChannel struct {
	subscribers sync.Map
}

func NewBroadcastChannel() *BroadcastChannel {
	return &BroadcastChannel{
		subscribers: sync.Map{},
	}
}

func (c *BroadcastChannel) Broadcast(v *Job) {
	value, ok := c.subscribers.Load(v.JobId)
	if !ok {
		log.Warn().Msgf("job id %s not exist", v.JobId)
		return
	}

	for _, ch := range value.([]chan *Job) {
		// send message to all listeners for this job ID
		ch <- v
		if v.Status == Canceled || v.Status == Complete || v.Status == Failed {
			// job is done running, no more updates
			close(ch)
			continue
		}
	}
}

func (c *BroadcastChannel) Subscribe(jobId string) chan *Job {
	channel := make(chan *Job, 1)
	val, ok := c.subscribers.Load(jobId)
	if ok {
		// existing jobid
		c.subscribers.Store(jobId, append(val.([]chan *Job), channel))
	} else {
		// new job id subscriber
		c.subscribers.Store(jobId, []chan *Job{channel})
	}

	return channel
}

func (c *BroadcastChannel) Unsubscribe(jobId string, ch chan *Job) {
	val, ok := c.subscribers.Load(jobId)
	if !ok {
		log.Warn().Msgf("No channels found for job id %s", jobId)
		return
	}

	tmp := val.([]chan *Job)
	for i, item := range tmp {
		if item == ch {
			// remove channel
			newList := append(tmp[:i], tmp[i+1:]...)
			if len(newList) == 0 {
				c.subscribers.Delete(jobId)
				return
			}

			c.subscribers.Store(jobId, newList)
			return
		}
	}
}
