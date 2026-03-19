// maps to java directory manager

package main

type DirectoryManager struct {
	files map[string]FileInfo
}

func NewDirectoryManager() *DirectoryManager {
	return &DirectoryManager {
		files: make(map[string]FileInfo),
	}
}

// enter inserts or updates file metadata
func (dm *DirectoryManager) Enter(name string, info FileInfo) {
	dm.files[name] = info
}

// returns fileinfo and whether it exists
func (dm *DirectoryManager) Lookup(name string) (FileInfo, bool) {
	info, valid := dm.files[name]
	return info, valid
}