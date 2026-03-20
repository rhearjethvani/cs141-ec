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

func (pm *PrinterManager) Count() int {
	return pm.resourceManager.Count()
}

// pass through methods
func (pm *PrinterManager) Request() int {
	return pm.resourceManager.Request()
}

func (pm *PrinterManager) Release(index int) {
	pm.resourceManager.Release(index)
}