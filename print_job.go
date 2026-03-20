// maps to java printjobthread

package main

import (
	"fmt"
)

type PrintJob struct {
	FileName string
}

func NewPrintJob(fileName string) *PrintJob {
	return &PrintJob {
		FileName: fileName,
	}
}

// executes the print job using the provided directory, disks, and priinter
func (pj *PrintJob) Run(directory *DirectoryManager, disks []*Disk, printers []*Printer, printerManager *PrinterManager) {
	info, valid := directory.Lookup(pj.FileName)
	if !valid {
		fmt.Println("File not found in directory:", pj.FileName)
		return
	}

	printerNum := printerManager.Request()
	if printerNum == -1 {
		fmt.Println("No printer available for file:", pj.FileName)
		return
	}

	for i := 0; i < info.FileLength; i++ {
		sector := info.StartingSector + i
		data := disks[info.DiskNumber].Read(sector)
		printers[printerNum].PrintLine(data)

		fmt.Println("Printed line from disk sector", sector, ":", data)
	}

	printerManager.Release(printerNum)
}