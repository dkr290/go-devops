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

	//create ticker and pass it directly
	ticker := time.NewTicker(1 * time.Second)
	go tickerCounter(ticker, c)
	time.Sleep(5 * time.Second)
	ticker.Stop()

	// Setting the value of channel
	c <- true
	time.Sleep(10 * time.Second)

	fmt.Println("Exitting")

}

func tickerCounter(ticker *time.Ticker, c chan bool) {

	// Using for loop
	for {

		// Select statement
		select {

		// Case statement
		case <-c:
			fmt.Println("Completed!")
			return

		// Case to print current time
		case tm := <-ticker.C:
			fmt.Println("The Current time is: ", tm)
		}
	}

}
