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