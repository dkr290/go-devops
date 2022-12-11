package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
)

func DeleteEntry(key string, csvfile string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}
	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	delete(index, key)

	err := saveCSVFile(csvfile)
	if err != nil {
		return err
	}
	return nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}
