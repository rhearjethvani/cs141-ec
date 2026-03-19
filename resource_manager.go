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