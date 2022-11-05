package main

import (
	"fmt"
	"time"
)

func compEqual() {
	now := time.Now()
	now_too := now
	time.Sleep(2 * time.Millisecond)
	later := time.Now()

	if now.Equal(now_too) {
		fmt.Printf("The two tiume variables are equal\n")
	} else {
		fmt.Printf("The two time variables are different \n")
	}

	if now.Equal(later) {
		fmt.Printf("The two time variables are equal\n")
	} else {
		fmt.Printf("The two time variables are different\n")
	}

}

func durationCalcul() {
	start := time.Now()
	fmt.Println("The script started...", start)
	sum := 0
	for i := 0; i < 100000000000; i++ {
		sum += 1
	}

	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("The script complerted at ", end)
	fmt.Println("The task took: ", duration.Hours(), "houurs to complete")
	fmt.Println("The task took: ", duration.Minutes(), "minutes to complete")
	fmt.Println("The task took: ", duration.Seconds(), "seconds to complete")
	fmt.Println("The task took: ", duration.Nanoseconds(), "nanoseconds to complete")

}

func transactionDeadline() {

	deadline_seconds := time.Duration((600 * 10) * time.Microsecond)
	Start := time.Now()
	fmt.Println("Deadline for the transaction is ", deadline_seconds)
	fmt.Println("The transactiopn has started", Start)

	sum := 0

	for i := 0; i < 2500000000; i++ {
		sum += 1
	}

	End := time.Now()
	Duration := End.Sub(Start)
	TrasactionTime := time.Duration(Duration.Nanoseconds()) * time.Nanosecond
	fmt.Println("The transaction is completed ", End, Duration)
	if TrasactionTime <= deadline_seconds {
		fmt.Println("Performance is OK transaction completed in ", TrasactionTime)
	} else {
		fmt.Println("Performance problem, transaction completed in", TrasactionTime, "seconds")
	}
}
