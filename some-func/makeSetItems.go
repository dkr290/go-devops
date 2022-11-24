package main

import "fmt"

type Set map[string]struct{}

func makeSetItems() {

	s := make(Set)

	s["item1"] = struct{}{}
	s["item2"] = struct{}{}

	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var r []string

	for k, _ := range s {
		r = append(r, k)
	}

	return r
}
