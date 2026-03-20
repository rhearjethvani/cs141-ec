package main

import "sync"

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