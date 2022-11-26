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

// other way is to define error

//var errFindMember = errors.New("The member was not fiund")

type findError struct {
	Name, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findMember(name, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	if v, ok := scMapping[name]; !ok {
		//return -1, errors.New("the member was not found in")
		return -1, findError{name, server, "The mumber was not found"}

	} else {
		return v, nil
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())
	if cl, err := findMember("Ruko", "Server 1"); err != nil {
		fmt.Println("Error occured while searching for clearence level: ", err)
		if v, ok := err.(findError); ok {
			fmt.Println("Server name is ", v.Server)
			fmt.Println("The member is ", v.Name)
		}
	} else {
		fmt.Println("The member was found: ", cl)
	}

}
