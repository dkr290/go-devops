package main

import "fmt"

func main() {

	a := []int8{99, 2, 54, 7, 3, 34, 123, 89}
	b := []int64{44, 6, 7, 293, 553, 7, 99, 8, 438, 3}

	sortInts(a)
	sortInts(b)

	fmt.Println(a)
	fmt.Println(b)

}

func sortInts[I int8 | int16 | int32 | int64](s []I) {
	sorted := false

	for !sorted {
		sorted = true
		for i := range s[:len(s)-1] {
			if s[i] > s[i+1] {
				sorted = false
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}

}
