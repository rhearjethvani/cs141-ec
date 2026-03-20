// maps to java disk manager

package main

type DiskManager struct {
	resourceManager *ResourceManager
	nextFreeSector []int
}

func NewDiskManager(numDisks int) *DiskManager {
	return &DiskManager {
		resourceManager: NewResourceManager(numDisks),
		nextFreeSector: make([]int, numDisks),
	}
}

// returns the next available sector index for a disk
func (dm *DiskManager) GetNextFreeSector(disk int) int {
	return dm.nextFreeSector[disk]
}

// update the next available sector index for a disk
func (dm *DiskManager) SetNextFreeSector(disk int, nextSector int) {
	dm.nextFreeSector[disk] = nextSector
}

// pass through methods
func (dm *DiskManager) Request() int {
	return dm.resourceManager.Request()
}

func (dm *DiskManager) Release(index int) {
	dm.resourceManager.Release(index)
}