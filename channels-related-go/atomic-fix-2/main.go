package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	//shutdown is flag to alert running goroutines to shutdown
	shutdown int64

	//wait for program to finish
	wg sync.WaitGroup
)

func main() {

	wg.Add(3)
	//create three go routines

	go doWork("A")
	go doWork("B")
	go doWork("C")

	//give gouroutines time to run
	time.Sleep(1 * time.Second)
	//safetly flag it is time to shutdown

	fmt.Println("Shutdown now")
	atomic.StoreInt64(&shutdown, 1)

	//wait for goroutines to finish
	wg.Wait()

}

// doWork simulates a goroutine performing work and
// checking the Shutdown flag to terminate early.

func doWork(name string) {

	defer wg.Done()

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		//Do we need to shutdown.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
