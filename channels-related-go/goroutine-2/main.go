package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(2)
	wg.Add(2)

	//create the gorotines

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	//wait the gorotine to finish
	fmt.Println("waiting to finish")
	wg.Wait()

	fmt.Println("Terminating the program")

}

func printPrime(prefix string) {

	// defert when its done dto decrease counter
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)

	}
	fmt.Println("Completed", prefix)
}
