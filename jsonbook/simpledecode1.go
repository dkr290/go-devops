package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Greeting struct {
	SomeMessage string `json:"message"`
}

func unmarsh1() {

	data := []byte(`{  
	"message": "Greeting from golang unmarshal test"
	}`)

	if !json.Valid(data) {
		fmt.Printf("JSON is not valid %s", data)
		os.Exit(1)
	}
	var v Greeting
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("The unmarshaled data is: ", v.SomeMessage)

}
