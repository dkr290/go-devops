package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	check := checkArgs(os.Args)

	myMp := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}

	if k, ok := myMp[check]; ok {
		fmt.Println("Found:", k)
	} else {
		log.Fatalln("The key not found in the map")
	}

}

func checkArgs(args []string) string {

	for i := 0; i < len(args); i++ {
		if len(args) != 2 {
			log.Fatalln("you did not pass one argument")
		}
	}

	return args[1]
}
