package main

import (
	"fmt"
	"os"
	"strconv"
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
}