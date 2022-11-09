package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	only_after, _ := time.Parse(time.RFC3339, "2020-11-01T22:08:41+00:00")
	fmt.Println(now, only_after)
	fmt.Println(now.After(only_after))

	if now.After(only_after) {
		fmt.Println("Executing actions")
	} else {
		fmt.Println("now is not the time yet")
	}

	fmt.Println("")

	compEqual()
	//durationCalcul()
	transactionDeadline()
	addTime()

	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Println(elapsedTime(start, end))

	fmt.Println(timeDiff("America/Los_Angeles"))

}
