package models

import (
	"sync"
	"sync/atomic"
)

type CountingMutex struct {
	mu           sync.Mutex
	waitingCount int32 // Atomic counter for waiting goroutines
}

func NewCountMutex() *CountingMutex {
	return &CountingMutex{
		waitingCount: 0,
		mu:           sync.Mutex{},
	}
}

func (wm *CountingMutex) Lock() {
	// Increment waiting count before attempting to acquire lock
	atomic.AddInt32(&wm.waitingCount, 1)
	wm.mu.Lock()
	// Decrement waiting count after acquiring lock
	atomic.AddInt32(&wm.waitingCount, -1)
}

func (wm *CountingMutex) Unlock() {
	wm.mu.Unlock()
}

func (wm *CountingMutex) WaitingCount() int32 {
	return atomic.LoadInt32(&wm.waitingCount)
}
