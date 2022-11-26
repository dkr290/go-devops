package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simulated database like data
var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

func findMember(name, server string) int {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	if v, ok := scMapping[name]; !ok {

		panic("The member was not found")

	} else {
		return v
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Apanic occured", err)
		}
	}()

	cl := findMember("Rsdsdggh", "Server 1")

	fmt.Println("The member was found: ", cl)

}
