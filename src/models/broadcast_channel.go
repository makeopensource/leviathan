package models

import (
	"context"
	"github.com/rs/zerolog/log"
)

type BroadcastChannel struct {
	subscribers Map[string, chan *Job]
}

func NewBroadcastChannel() (*BroadcastChannel, context.Context) {
	bc := &BroadcastChannel{
		subscribers: Map[string, chan *Job]{},
	}
	return bc, context.WithValue(context.Background(), "broadcast", bc)
}

func (c *BroadcastChannel) Broadcast(v *Job) {
	ch, ok := c.subscribers.Load(v.JobId)
	if !ok {
		log.Warn().Msgf("job update channel %s does not exist", v.JobId)
		return
	}

	ch <- v
	if v.Status.Done() {
		// job is done running, no more updates
		close(ch)
	}
}

func (c *BroadcastChannel) Subscribe(jobId string) chan *Job {
	channel := make(chan *Job, 2)
	val, ok := c.subscribers.Load(jobId)
	// if exists close old channel
	if ok {
		close(val)
	}

	// new job id subscriber
	c.subscribers.Store(jobId, channel)
	return channel
}

func (c *BroadcastChannel) Unsubscribe(jobId string) {
	c.subscribers.Delete(jobId)
}
