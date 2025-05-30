package sync_utils

type WorkerSemaphore struct {
	semaphore chan struct{}
}

func NewWorkerSemaphore(maxWorkers int) *WorkerSemaphore {
	sem := &WorkerSemaphore{
		semaphore: make(chan struct{}, maxWorkers),
	}
	for i := 0; i < maxWorkers; i++ {
		sem.Release() // fill the semaphore with maxWorkers
	}
	return sem
}

// Acquire returns the semaphore channel for select statements
// will block until resource is available
func (s *WorkerSemaphore) Acquire() <-chan struct{} {
	return s.semaphore
}

func (s *WorkerSemaphore) Release() {
	s.semaphore <- struct{}{}
}
