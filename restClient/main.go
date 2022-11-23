package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	resr, err := http.Get("http://quotes.rest/qod.json")
	if err != nil {
		log.Fatalln("Error of get", err)
	}

	defer resr.Body.Close()

	body, err := io.ReadAll(resr.Body)
	if err != nil {
		log.Fatalln("Error reading body", err)
	}

	fmt.Println(string(body))

}
