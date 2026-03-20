// maps to java printjobthread

package main

type PrintJob struct {
	FileName string
}

func NewPrintJob(fileName string) *PrintJob {
	return &PrintJob {
		FileName: fileName,
	}
}