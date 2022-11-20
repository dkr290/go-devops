package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func readCsv() {

	in := `firstName,lastName,age
	       Celina, Jones, 19
		   Calyan, Reddy,29
		   John,Nalm,89
	`

	r := csv.NewReader(strings.NewReader(in))
	header := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if !header {
			for idx, value := range record {
				switch idx {
				case 0:
					fmt.Println("First name is:", value)
				case 1:
					fmt.Println("The last name is:", value)
				case 2:
					fmt.Println("The age is:", value)
				}

			}
		}
		header = false
	}
}
