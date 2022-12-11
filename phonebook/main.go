package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/dkr290/go-devops/phonebook/helpers"
)

// CSVFILE resides in the home directory of the current user
var CSVFILE = "./csv.data"

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s insert|delete|search|list <arguments>\n", exe)
		return
	}

	_, err := os.Stat(CSVFILE)

	// meaning error is not nil so the file does not exists
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			log.Fatalln("Cannot create the file", err)

		}

		f.Close()
	}

	fileinfo, _ := os.Stat(CSVFILE)
	// is it regular file
	mode := fileinfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "not a regular file")
		return
	}
	err = helpers.ReadCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = helpers.CreateIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	// Differentiating between the commands
	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !helpers.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := helpers.Inits(arguments[2], arguments[3], t)
		// If it was nil, there was an error
		if temp != nil {
			err := helpers.Insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !helpers.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		err := helpers.DeleteEntry(t, CSVFILE)
		if err != nil {
			fmt.Println(err)
		}

	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !helpers.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := helpers.Search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(*temp)
	case "list":
		helpers.List()
		// if nothing has been match for command line argument
	default:
		fmt.Println("Not valid option")

	}
}
