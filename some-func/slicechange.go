package main

import "fmt"

func slice1() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a = append(a[:2], a[4:]...)
	fmt.Println(a)
}
