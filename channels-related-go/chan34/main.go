package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numofGr    = 4  //num of go routines
	tasktoLoad = 10 //amount of work to process
)

var wg sync.WaitGroup

func init() {

	//seed random number
	rand.Seed(time.Now().Unix())
}

func main() {

	//create buffered channel to manage the task load

	tasks := make(chan string, tasktoLoad)

	// Launch goroutines to handle the work.

	wg.Add(numofGr)

	for gr := 1; gr <= numofGr; gr++ {
		go worker(tasks, gr)
	}

	//add bunch of work to get done
	for post := 1; post <= tasktoLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// Close the channel so the goroutines will quit when all work finish

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {

	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//this means the channel is empty or closed
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}
		// Display we are starting the work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		// Randomly wait to simulate work time.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display we finished the work.
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)

	}

}
