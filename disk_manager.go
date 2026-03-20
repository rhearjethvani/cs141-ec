// maps to java disk manager

package main

import "sync"

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