package main

import "fmt"

const (
	Winter = 12 - (iota + 3)
	Spring = 3
	Summer = 6
	Fall   = 9
)

func main() {

	fmt.Println(Winter, Spring, Summer, Fall)

}
