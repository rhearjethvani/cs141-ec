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