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

// returns an available resource index; for user0 return first free one
func (rm *ResourceManager) Request() int {
	for i := 0; i < len(rm.isFree); i++ {
		if rm.isFree[i] {
			rm.isFree[i] = false
			return i
		}
	}

	// if none available
	return -1
}

// frees a resource index
func (rm *ResourceManager) Release(index int) {
	rm.isFree[index] = true
}