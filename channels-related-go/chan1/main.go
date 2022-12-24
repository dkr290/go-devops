package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// create uinbuffered channel

	court := make(chan int)

	//add count of the two for each go routine

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	//start the set
	court <- 1

	//wait for game to finish

	wg.Wait()

}

//player simulates a person playing game of tennis

func player(name string, court chan int) {
	//schedule the call to Done to tell the main we are done

	defer wg.Done()

	for {
		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {
			//If the channel was closed we won
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// pick a random number and see if we miss the ball
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// Hit the ball back to the opposing player.
		court <- ball
	}
}
