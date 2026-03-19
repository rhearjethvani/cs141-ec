// maps to java printer

package main

const PrintDelayMs = 2750

type Printer struct {
	ID int
}

func NewPrinter(id int) *Printer {
	return &Printer {
		ID: id,
	}
}