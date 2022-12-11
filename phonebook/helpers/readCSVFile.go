package helpers

import (
	"encoding/csv"
	"os"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var data = []Entry{}

func ReadCSVFile(filepath string) error {

	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	// read all at once in the csv file

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		data = append(data, temp)
	}

	return nil

}
