package main

import (
	"fmt"
	"sync"
	"time"
)

// wg is used to wait for the program to finis
var wg sync.WaitGroup

func main() {

	//create new unbuffered channel
	baton := make(chan int)

	//add count of one for count of the last runner

	wg.Add(1)
	//first runner to hit mark
	go Runner(baton)

	//Start the race
	baton <- 1

	wg.Wait()

}

// Runner simulates a person running in the relay race.

func Runner(baton chan int) {

	var newRunner int
	//wait to receive baton

	runner := <-baton

	//start the runner around the track
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// New runner to the line.
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}
	// Running around the track.
	time.Sleep(100 * time.Millisecond)

	// Is the race over.
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)
	baton <- newRunner
}
