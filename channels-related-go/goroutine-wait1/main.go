package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting the goroutines")

	//declare annonimout function and create goroutine

	go func() {
		//schedule the call to Done to tell main we are done
		defer wg.Done()
		//display alphabet three times
		for count := 0; count <= 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//wait for goroutine to finish
	fmt.Println("waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating the program")

}
