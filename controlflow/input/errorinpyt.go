package main

import (
	"fmt"
	"strconv"
)

func errInputType() {

	args := returnArgs()

	var total, nInts, nFloats int

	invalid := make([]string, 0)

	for _, k := range args[1:] {
		//check if it is an integer
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue

		}

		// is it a float typre

		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		//then maybe is invalid

		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)
	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
