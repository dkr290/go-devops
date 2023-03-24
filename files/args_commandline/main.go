package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	fmt.Println(args)

	var iRags = []int{}
	for _, i := range args {
		v, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iRags = append(iRags, v)
	}

	max := 0

	for _, v := range iRags {
		if v > max {
			max = v
		}
	}

	fmt.Println("max value :", max)

}
