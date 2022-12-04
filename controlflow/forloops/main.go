package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}

	i := 0

	// this way is not commonly used (do-while approach)
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}

	fmt.Println()

	//simulating while loop now

	i = 0

	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// This is a slice but range also works with arrays
	aSlice := []int{-1, 2, 1, -1, 2, -2}

	for i, v := range aSlice {
		fmt.Println("index: ", i, "value: ", v)
	}

}
