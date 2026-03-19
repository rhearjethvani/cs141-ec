// maps to java fileinfo

package main

import "fmt"

type FileInfo struct {
	DiskNumber int
	StartingSector int
	FileLength int
}

func NewFileInfo(disk int, start int, length int) FileInfo {
	return FileInfo {
		DiskNumber: disk,
		StartingSector: start,
		FileLength: length,
	}
}

func (f FileInfo) String() string {
	return fmt.Sprintf("[disk=%d start=%d len=%d]", f.DiskNumber, f.StartingSector, f.FileLength)
}