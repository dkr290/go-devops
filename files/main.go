package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//cleanUp()
	//checkFiles("junk.txt", "test.txt")

	CreateFile("text.txt")
	RAll("text.txt")

	backupFile := "backupFile1.txt"
	workingFile := "note.txt"
	data := "This is second note"

	err := createBackup(workingFile, backupFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		note := data + " " + strconv.Itoa(i)
		err := addNotes(workingFile, note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println("Read csv data")
	readCsv()

}
