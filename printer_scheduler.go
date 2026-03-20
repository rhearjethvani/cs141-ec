package main

import (
	"sync"
	"fmt"
)

func StartPrinterScheduler(
	printQueue *PrintQueue,
	disks []*Disk,
	printers []*Printer,
	printerManager *PrinterManager,
	printWG *sync.WaitGroup,
) {
	go func() {
		for {
			job := printQueue.Dequeue()

			fmt.Println("Scheduling print job:", job.FileName, "priority:", job.Priority)

			printerNum := printerManager.Request()
			printer := printers[printerNum]

			for i := 0; i < job.Info.FileLength; i++ {
				sector := job.Info.StartingSector + i
				data := disks[job.Info.DiskNumber].Read(sector)
				printer.PrintLine(data)
			}

			printerManager.Release(printerNum)
			printWG.Done()
		}
	}()
}