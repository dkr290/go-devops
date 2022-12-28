package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	filename := "words.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(data))

	fmt.Printf("fount %d words", len(words))

}
