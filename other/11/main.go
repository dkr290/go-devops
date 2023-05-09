package main

import "fmt"

func main() {
	xVar := 1

	for true {
		if xVar == 5 {
			break
		}
		fmt.Println(xVar)
		xVar++
	}
}
