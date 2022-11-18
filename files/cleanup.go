package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanUp() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		for {
			s := <-sigs

			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might might of pressed CTRL-C")
				fmt.Println("Some cleanup occured")
				cleanUpf()
				done <- true
			case syscall.SIGTERM:
				fmt.Println()
				fmt.Println("someone passed CTRL-Z")
				fmt.Println("Some cleanup occured")
				cleanUpf()
				done <- true

			}
		}
	}()

	fmt.Println("Program is blocked untils signal is caught CTRL-z or CTRL-C")
	<-done
	fmt.Println("Quit")

}

func cleanUpf() {
	fmt.Println("Simulating clean up")
	for i := 0; i <= 10; i++ {
		fmt.Println("Deleting files... Not Really", i)
		time.Sleep(1 * time.Second)
	}

}
