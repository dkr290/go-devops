package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("One\n")
	c := make(chan bool)

	go testFunc(c)
	fmt.Printf("two\n")

	areWeFinished := <-c
	fmt.Printf("areWeFinished %v\n", areWeFinished)

	fmt.Println("")
	newc := make(chan bool)
	fmt.Println("Printing from main")
	go printSomething(newc, "Hello from inside the channel passed func")
	newc <- true
	<-newc

}

func testFunc(c chan bool) {

	for i := 0; i < 5; i++ {
		fmt.Printf("checking...\n")
		time.Sleep(1 * time.Second)
	}

	c <- true

}

func printSomething(c chan bool, s string) {

	if b := <-c; b {
		fmt.Println("THis is a simple test print: ", s)
	}
	c <- true
}
