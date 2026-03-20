// maps to java resource manager

package main

import (
	"sync"
)

type ResourceManager struct {
	isFree []bool
	mu sync.Mutex
	cond *sync.Cond
}

func NewResourceManager(count int) *ResourceManager {
	rm := &ResourceManager {
		isFree: make([]bool, count),
	}

	for i := 0; i < count; i++ {
		rm.isFree[i] = true
	}

	rm.cond = sync.NewCond(&rm.mu)
	return rm
}

// return number of resources
func (rm *ResourceManager) Count() int {
	return len(rm.isFree)
}

// is a specific resource available
func (rm *ResourceManager) IsFree(index int) bool {
	return rm.isFree[index]
}

// returns an available resource index; for user0 return first free one
func (rm *ResourceManager) Request() int {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	for {
		for i := 0; i < len(rm.isFree); i++ {
			if rm.isFree[i] {
				rm.isFree[i] = false
				return i
			}
		}

		rm.cond.Wait()
	}
}

// frees a resource index
func (rm *ResourceManager) Release(index int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.isFree[index] = true
	rm.cond.Broadcast()
}