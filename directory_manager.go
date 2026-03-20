// maps to java directory manager

package main

import "sync"

type DirectoryManager struct {
	files map[string]FileInfo
	mu sync.Mutex
}

func NewDirectoryManager() *DirectoryManager {
	return &DirectoryManager {
		files: make(map[string]FileInfo),
	}
}

// enter inserts or updates file metadata
func (dm *DirectoryManager) Enter(name string, info FileInfo) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	dm.files[name] = info
}

// returns fileinfo and whether it exists
func (dm *DirectoryManager) Lookup(name string) (FileInfo, bool) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	
	info, valid := dm.files[name]
	return info, valid
}