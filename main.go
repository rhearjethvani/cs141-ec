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
	
	diskNum := 0
	startSector := 0
	fileLength := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, ".save") {
			saving = true
			currentFileName = strings.TrimSpace(line[len(".save"):])
			diskNum = 0
			startSector = diskManager.GetNextFreeSector(diskNum)
			fileLength = 0

			fmt.Println("SAVE command for file:", currentFileName)
			fmt.Println("Starting save on disk", diskNum, "at sector", startSector)
		} else if line == ".end" {
			fmt.Println("END command for file:", currentFileName)

			if saving {
				info := NewFileInfo(diskNum, startSector, fileLength)
				directory.Enter(currentFileName, info)
				diskManager.SetNextFreeSector(diskNum, startSector+fileLength)

				fmt.Println("Saved file metadata:", info)
			}

			saving = false
			currentFileName = ""
			fileLength = 0
			startSector = 0
		} else if strings.HasPrefix(line, ".print") {
			fileNameToPrint := strings.TrimSpace(line[len(".print"):])
			fmt.Println("PRINT command for file:", fileNameToPrint)

			job := NewPrintJob(fileNameToPrint)
			job.Run(directory, disks, printers[0])
		} else if saving {
			targetSector := startSector + fileLength
			disks[diskNum].Write(targetSector, line)
			fileLength++

			fmt.Println("Wrote data line to disk sector", targetSector, ":", line)
		} else {
			fmt.Println("Ignoring unexpected line:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	file.Close()
}