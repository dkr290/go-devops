package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	start := time.Now()

	if len(os.Args != 2) {

		fmt.Println("Usage: dates_parse_string")
		return
	}

	dateStr := os.Args[1]

	d, err := time.Parse("02 Jan 2006", dateStr)
	if err != nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

}
