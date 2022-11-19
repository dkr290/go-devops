package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	ErrorWorkingFileNotFound = errors.New("The working file was not found")
)

func createBackup(working, backup string) error {

	_, err := os.Stat(working)
	if err != nil {
		if !os.IsNotExist(err) {
			return ErrorWorkingFileNotFound
		}

		return err
	}

	workingFile, err := os.Open(working)
	if err != nil {
		return err
	}

	fileContents, err := ioutil.ReadAll(workingFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(backup, fileContents, 0644)
	if err != nil {
		return fmt.Errorf("Could not write to the backup file %s ", err)
	}

	return nil

}

func addNotes(workingFile, notes string) error {

	notes += "\n"

	f, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write([]byte(notes)); err != nil {
		return err
	}

	return nil
}
