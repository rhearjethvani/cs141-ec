package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . <numUsers> <numDisks> <numPrinters>")
		return
	}

	numUsers, err1 := strconv.Atoi(os.Args[1])
	numDisks, err2 := strconv.Atoi(os.Args[2])
	numPrinters, err3 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("All arguments must be integers")
		return
	}

	fmt.Println("141OS Configuration:")
	fmt.Println("Users:", numUsers)
	fmt.Println("Disks:", numDisks)
	fmt.Println("Printers:", numPrinters)

	// initialize runtime components for single-user testing
	fmt.Println("Initialized runtime:")

	disks := make([]*Disk, numDisks)
	for i := 0; i < numDisks; i++ {
		disks[i] = NewDisk(i)
	}
	fmt.Println("Disks created:", len(disks))

	printers := make([]*Printer, numPrinters)
	for i := 0; i < numPrinters; i++ {
		printers[i] = NewPrinter(i)
	}
	fmt.Println("Printers created:", len(printers))

	directory := NewDirectoryManager()
	fmt.Println("Directory ready:", directory != nil)

	diskManager := NewDiskManager(numDisks)
	fmt.Println("DiskManager ready:", diskManager != nil)

	printerManager := NewPrinterManager(numPrinters)
	fmt.Println("PrinterManager ready:", printerManager != nil)

	// smoke test: read USER0
	user := NewUser(0)
	path := user.InputPath()

	fmt.Println("Reading from:", path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	saving := false
	currentFileName := ""

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, ".save") {
			saving = true
			currentFileName = strings.TrimSpace(line[len(".save"):])
			fmt.Println("SAVE command for file:", currentFileName)
		} else if line == ".end" {
			fmt.Println("END command for file:", currentFileName)
			saving = false
			currentFileName = ""
		} else if strings.HasPrefix(line, ".print") {
			fileNameToPrint := strings.TrimSpace(line[len(".print"):])
			fmt.Println("PRINT command for file:", fileNameToPrint)
		} else if saving {
			fmt.Println("DATA line:", line)
		} else {
			fmt.Println("Ignoring unexpected line:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	file.Close()
}