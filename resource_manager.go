// maps to java resource manager

package main

type ResourceManager struct {
	isFree []bool
}

func NewResourceManager(count int) *ResourceManager {
	isFree := make([]bool, count)
	for i := range isFree {
		isFree[i] = true
	}

	return &ResourceManager {
		isFree: isFree,
	}
}

// return number of resources
func (rm *ResourceManager) Count() int {
	return len(rm.isFree)
}

// is a specific resource available
func (rm *ResourceManager) IsFree(index int) bool {
	return rm.isFree[index]
}