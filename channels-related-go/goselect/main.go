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

func findMember(name, server string, c chan int) {
	//simulate searching
	time.Sleep(time.Duration(rand.Intn(50)) * time.Minute)

	c <- scMapping[name]
}

func main() {

	rand.Seed(time.Now().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)

	name := "James"
	go findMember(name, "Server 1", c1)
	go findMember(name, "Server 2", c2)

	select {
	case sc := <-c1:
		fmt.Println(name, " has been found with id ", sc, "on the server 1")
	case sc := <-c2:
		fmt.Println(name, "has been found with id ", sc, "on the serv 2")
	// case <-time.After(2 * time.Minute):
	// 	fmt.Println("Search timed out")
	// }
	default:
		fmt.Println("Too slow")

	}

}
