package main

import "sync"

type PrintQueue struct {
	queue []PrintJob
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewPrintQueue() *PrintQueue {
	pq := &PrintQueue{}
	pq.cond = sync.NewCond(&pq.mu)
	return pq
}

func (pq *PrintQueue) Enqueue(job PrintJob) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if job.Priority == 1 {
		inserted := false
		for i, existing := range pq.queue {
			if existing.Priority == 0 {
				pq.queue = append(pq.queue[:i], append([]PrintJob{job}, pq.queue[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			pq.queue = append(pq.queue, job)
		}
	} else {
		pq.queue = append(pq.queue, job)
	}

	pq.cond.Signal()
}

func (pq *PrintQueue) Dequeue() PrintJob {
	pq.mu.Lock()
	for len(pq.queue) == 0 {
		pq.cond.Wait()
	}
	job := pq.queue[0]
	pq.queue = pq.queue[1:]
	pq.mu.Unlock()
	return job
}