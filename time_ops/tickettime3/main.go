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

	//create ticker and pass it directly
	ticker := time.NewTicker(1 * time.Second)
	c := tickerCounter(ticker)
	time.Sleep(5 * time.Second)
	ticker.Stop()

	// Setting the value of channel
	c <- true

}

func tickerCounter(ticker *time.Ticker) chan bool {
	c := make(chan bool)
	go func() {
		// Using for loop
	LOOP:
		for {

			// Select statement
			select {

			// Case statement
			case <-c:
				fmt.Println("Completed!")
				break LOOP

			// Case to print current time
			case tm := <-ticker.C:
				fmt.Println("The Current time is: ", tm)
			}
		}
		fmt.Println("Exitting the tick counter...")
	}()
	return c

}
