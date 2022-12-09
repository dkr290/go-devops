package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type record struct {
	Field1 int
	Field2 string
}

func main() {

	s := []record{}

	for i := 0; i < 10; i++ {

		temp := record{Field1: i, Field2: "text " + strconv.Itoa(i)}
		s = append(s, temp)
	}

	// Accessing the fields of the first element
	fmt.Println("Index 0:", s[0].Field1, s[0].Field2)
	fmt.Println("Number of structures:", len(s))
	sum := 0
	for _, k := range s {
		sum += k.Field1
	}
	fmt.Println("Sum:", sum)

	resp := matchInt(strconv.Itoa(sum))

	fmt.Println(resp)
	n1 := matchNameSur("Somename")
	fmt.Println(n1)
}

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}
