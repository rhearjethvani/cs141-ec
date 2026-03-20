// maps to java printer

package main

import (
	"time"
	"fmt"
	"os"
)

const PrintDelayMs = 2 // 2750

type Printer struct {
	ID int
}

func NewPrinter(id int) *Printer {
	return &Printer {
		ID: id,
	}
}

func (p *Printer) PrintLine(data string) {
	time.Sleep(time.Millisecond * PrintDelayMs)

	filename := fmt.Sprintf("PRINTER%d", p.ID)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening printer file:", err)
		return
	}

	_, err = f.WriteString(data + "\n")
	if err != nil {
		fmt.Println("Error writing to printer file:", err)
	}

	f.Close()
}