// maps to java disk manager

package main

import (
	"sync"
	"fmt"
)

type FreeSegment struct {
	Start int
	Length int
}

type DiskManager struct {
	resourceManager *ResourceManager
	nextFreeSector []int
	freeSegments [][]FreeSegment
	mu sync.Mutex
}

func NewDiskManager(numDisks int) *DiskManager {
	return &DiskManager {
		resourceManager: NewResourceManager(numDisks),
		nextFreeSector: make([]int, numDisks),
		freeSegments: make([][]FreeSegment, numDisks),
	}
}

// returns the next available sector index for a disk
func (dm *DiskManager) GetNextFreeSector(disk int) int {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	return dm.nextFreeSector[disk]
}

// update the next available sector index for a disk
func (dm *DiskManager) SetNextFreeSector(disk int, nextSector int) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	
	dm.nextFreeSector[disk] = nextSector
}

// pass through methods
func (dm *DiskManager) Request() int {
	return dm.resourceManager.Request()
}

func (dm *DiskManager) Release(index int) {
	dm.resourceManager.Release(index)
}

// picks the best available disk for a new file of the given length
// prefers a disk with a reusable free segment; falls back to least-used disk
func (dm *DiskManager) ChooseDisk(length int) int {
	dm.resourceManager.mu.Lock()
	defer dm.resourceManager.mu.Unlock()

	for {
		// first pass: prefer a disk with a reusable segment that fits
		for i, free := range dm.resourceManager.isFree {
			if !free {
				continue
			}
			dm.mu.Lock()
			for _, seg := range dm.freeSegments[i] {
				if seg.Length >= length {
					dm.mu.Unlock()
					dm.resourceManager.isFree[i] = false
					fmt.Println("Chose disk", i, "(has reusable space)")
					return i
				}
			}
			dm.mu.Unlock()
		}

		// second pass: pick the free disk with the least data written
		best := -1
		bestSectors := int(^uint(0) >> 1)
		for i, free := range dm.resourceManager.isFree {
			if !free {
				continue
			}
			dm.mu.Lock()
			used := dm.nextFreeSector[i]
			dm.mu.Unlock()
			if used < bestSectors {
				bestSectors = used
				best = i
			}
		}

		if best != -1 {
			dm.resourceManager.isFree[best] = false
			fmt.Println("Chose disk", best, "(least used, sectors:", bestSectors, ")")
			return best
		}

		dm.resourceManager.cond.Wait()
	}
}

// tries to reuse a freed segment on this disk; if none fits, it appends at nextFreeSector
func (dm *DiskManager) AllocateSpace(disk int, length int) int {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	segments := dm.freeSegments[disk]
	for i, seg := range segments {
		if seg.Length >= length {
			start := seg.Start

			if seg.Length == length {
				dm.freeSegments[disk][i].Start += length
				dm.freeSegments[disk][i].Length -= length
			}

			return start
		}
	}

	start := dm.nextFreeSector[disk]
	dm.nextFreeSector[disk] += length
	return start
}

// marks a deleted file's disk region as reusable
func (dm *DiskManager) FreeSpace(disk int, start int, length int) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	dm.freeSegments[disk] = append(dm.freeSegments[disk], FreeSegment{
		Start: start,
		Length: length,
	})
}