// maps to java disk

package main

const NumSectors = 2048
const DiskDelayMs = 800

type Disk struct {
	ID int
	Sectors []string
}

func NewDisk(id int) *Disk {
	return &Disk {
		ID: id,
		Sectors: make([]string, NumSectors),
	}
}