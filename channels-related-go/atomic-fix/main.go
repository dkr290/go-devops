package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//counter is a varible accessed and incremented by all go routines

var counter int64

// wg is used to wait for the program to finish

var wg sync.WaitGroup

func main() {

	// add count of 2 for each go routine
	wg.Add(2)

	// creating 2 go routines

	go intCounter(1)
	go intCounter(2)

	//wait for go rotines to finish

	wg.Wait()

	//display the final value
	fmt.Println("Final counter is ", counter)

}

// incCounter increments the package level counter variable.

func intCounter(id int) {
	// after end decrement with done

	defer wg.Done()

	for count := 0; count < 2; count++ {
		//safely add 1 to the counter
		atomic.AddInt64(&counter, 1)
		// the thread to be placed back to the queue

		runtime.Gosched()
	}

}
