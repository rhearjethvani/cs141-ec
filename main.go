package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
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

	users := make([]*User, numUsers)
	for i :=  0; i < numUsers; i++ {
		users[i] = NewUser(i)
	}

	var userWG sync.WaitGroup
	var printWG sync.WaitGroup

	for _, user := range users {
		userWG.Add(1)
		
		go func(u *User) {
			defer userWG.Done()
			u.Run(disks, printers, directory, diskManager, printerManager, &printWG)
		} (user)
	}

	userWG.Wait()
	printWG.Wait()
}