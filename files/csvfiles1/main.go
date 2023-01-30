package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type record []string

func main() {

	fmt.Println(readRec())
	recs := []record{
		{"Miroslav", "Block"},
		{"Sarah", "Murphy"},
		{"David", "Justice"},
	}

	if err := writeRec(recs); err != nil {
		panic(err)
	}

}

func (r record) validate() error {
	if len(r) != 2 {
		return errors.New("data format is incorrect")
	}
	return nil

}

func (r record) first() string {

	return r[0]
}

func (r record) last() string {
	return r[1]
}

func readRec() ([]record, error) {
	b, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}

	defer b.Close()

	scanner := bufio.NewScanner(b)

	var records []record
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		var rec record = strings.Split(line, ",")
		if err := rec.validate(); err != nil {
			return nil, fmt.Errorf("entry at line %d was invalid: %w", lineNum, err)
		}
		records = append(records, rec)
		lineNum++
	}

	return records, nil

}

func writeRec(rec []record) error {
	file, err := os.OpenFile("data-sorted.csv", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	//sort by last name

	sort.Slice(
		rec,
		func(i, j int) bool {
			return rec[i].last() < rec[j].last()
		},
	)

	for _, re := range rec {
		_, err := file.Write(re.csv())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r record) csv() []byte {
	b := bytes.Buffer{}
	for _, field := range r {
		b.WriteString(field + ",")
	}
	return b.Bytes()
}
