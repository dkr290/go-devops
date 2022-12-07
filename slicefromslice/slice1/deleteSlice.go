package main

import "fmt"

func deleteElement(s []int, i int) {

	fmt.Println("Original slice:", s)

	// Delete element at index i
	if i > len(s)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	// The ... operator auto expands aSlice[i+1:] so that
	// its elements can be appended to aSlice[:i] one by one
	s = append(s[:i], s[i+1:]...)
	fmt.Println("After 1st deletion:", s)

	// Delete element at index i
	if i > len(s)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	// Replace element at index i with last element
	s[i] = s[len(s)-1]
	// Remove last element
	s = s[:len(s)-1]
	fmt.Println("After 2nd deletion:", s)
}
