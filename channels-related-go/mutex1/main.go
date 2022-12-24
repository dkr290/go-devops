package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (

	//variable incremented by all goroutines
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {

	//add count of two for the two go routines
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("The incremented counter ", counter)

}

// incCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.

func intCounter(id int) {

	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this
		// critical section at a time.

		mutex.Lock()
		{

			//capture the value of counter
			value := counter

			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}

}
