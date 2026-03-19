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