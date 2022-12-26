package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	LastName  string  `json:"lname"`
	FirstName string  `json:"fname"`
	Address   Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func marshal1() {

	p := Person{LastName: "Bonf", FirstName: "James"}
	p.Address.Street = "Galaxy street far away"
	p.Address.City = "Orange Side"
	p.Address.State = "Oklahoma"
	p.Address.ZipCode = 1234

	noPrettyPrint, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	prettyPrint, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(noPrettyPrint))

	fmt.Println("")
	fmt.Println(string(prettyPrint))

}
