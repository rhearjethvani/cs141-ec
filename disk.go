// maps to java disk

package main

import "time"

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

// writes data to a specific sector with delay
func (d *Disk) Write(sector int, data string) {
	time.Sleep(time.Millisecond * DiskDelayMs)
	d.Sectors[sector] = data
}

// copies data from a specific sector into a string result with delay
func (d *Disk) Read(sector int) string {
	time.Sleep(time.Millisecond * DiskDelayMs)
	return d.Sectors[sector]
}