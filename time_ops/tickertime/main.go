package main

// Importing fmt and time
import (
	"fmt"
	"time"
)

// Calling main
func main() {

	// Creating channel using make
	// keyword
	c := make(chan bool)

	go tickerCounter(2, c)
	// Setting the value of channel

	time.Sleep(7 * time.Second)
	c <- true

}

func tickerCounter(n int, c chan bool) {
	// Calling NewTicker method
	d := time.NewTicker(time.Duration(n) * time.Second)

	// Using for loop
	for {

		// Select statement
		select {

		// Case statement
		case <-c:
			fmt.Println("Completed!")
			return

		// Case to print current time
		case tm := <-d.C:
			fmt.Println("The Current time is: ", tm)
		}
	}

}
