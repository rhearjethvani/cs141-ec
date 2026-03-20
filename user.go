// maps to java userthread

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID int
}

func NewUser(id int) *User {
	return &User{
		ID: id,
	}
}

// filename the user reads from (ex. USER0)
func (u *User) InputFile() string {
	return "USER" + strconv.Itoa(u.ID)
}

// filepath resolution
func (u *User) InputPath() string {
	return filepath.Join("users", u.InputFile())
}

// processes user's command file
func (u *User) Run(
	disks []*Disk,
	printers []*Printer,
	directory *DirectoryManager,
	diskManager *DiskManager,
	printerManager *PrinterManager,
	printWG *sync.WaitGroup,
	printQueue *PrintQueue,
) {
	_ = printers
	_ = printerManager

	path := u.InputPath()
	fmt.Println("Reading from:", path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	saving := false
	currentFileName := ""

	diskNum := 0
	startSector := 0
	fileLines := make([]string, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, ".save") {
			saving = true
			currentFileName = strings.TrimSpace(line[len(".save"):])

			diskNum = diskManager.Request()
			fileLines = make([]string, 0)

			fmt.Println("SAVE command for file:", currentFileName)

		} else if line == ".end" {
			fmt.Println("END command for file:", currentFileName)

			if saving {
				fileLength := len(fileLines)
				startSector = diskManager.AllocateSpace(diskNum, fileLength)

				fmt.Println("Starting save on disk", diskNum, "at sector", startSector)

				for i, dataLine := range fileLines {
					targetSector := startSector + i
					disks[diskNum].Write(targetSector, dataLine)
					fmt.Println("Wrote data line to disk sector", targetSector, ":", dataLine)
				}

				info := NewFileInfo(diskNum, startSector, fileLength)
				directory.Enter(currentFileName, info)
				diskManager.Release(diskNum)

				fmt.Println("Saved file metadata:", info)
			}

			saving = false
			currentFileName = ""
			fileLines = make([]string, 0)

		} else if strings.HasPrefix(line, ".print") {
			fileNameToPrint := strings.TrimSpace(line[len(".print"):])
			fmt.Println("PRINT command for file:", fileNameToPrint)

			info, valid := directory.Lookup(fileNameToPrint)
			if !valid {
				fmt.Println("File not found in directory:", fileNameToPrint)
				continue
			}

			job := NewPrintJob(fileNameToPrint, info)
			printWG.Add(1)
			printQueue.Enqueue(job)

		} else if strings.HasPrefix(line, ".delete") {
			fileNameToDelete := strings.TrimSpace(line[len(".delete"):])
			fmt.Println("DELETE command for file:", fileNameToDelete)

			info, valid := directory.Lookup(fileNameToDelete)
			if !valid {
				fmt.Println("File not found for deletion:", fileNameToDelete)
				continue
			}

			deleted := directory.Delete(fileNameToDelete)
			if deleted {
				diskManager.FreeSpace(info.DiskNumber, info.StartingSector, info.FileLength)
				fmt.Println("Deleted file metadata for:", fileNameToDelete)
				fmt.Println(
					"Freed disk space on disk",
					info.DiskNumber,
					"start",
					info.StartingSector,
					"length",
					info.FileLength,
				)
			}

		} else if saving {
			fileLines = append(fileLines, line)
			fmt.Println("Buffered data line:", line)

		} else if line != "" {
			fmt.Println("Ignoring unexpected line:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}