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
}

func testFunc(c chan bool) {

	for i := 0; i < 5; i++ {
		fmt.Printf("checking...\n")
		time.Sleep(1 * time.Second)
	}

	c <- true

}
