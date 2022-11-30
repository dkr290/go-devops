package main

import "fmt"

func main() {
	itemsSold()

}

func itemsSold() {
	items := map[string]int{
		"John":    41,
		"Celine1": 109,
		"Micah":   28,
	}

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)

		if v < 40 {
			fmt.Println("is below expecations")
		} else if v > 40 && v <= 100 {
			fmt.Println("meet the expectations")
		} else if v > 100 {
			fmt.Println("exc expectations")
		}
	}
}
