// maps to java printer manager

package main

type PrinterManager struct {
	resourceManager *ResourceManager
}

func NewPrinterManager(numPrinters int) *PrinterManager {
	return &PrinterManager {
		resourceManager: NewResourceManager(numPrinters),
	}
}