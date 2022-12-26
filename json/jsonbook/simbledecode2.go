package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	StudentId  int      `json:"id"`
	LastName   string   `json:"lname"`
	MiddleName string   `json:"minitial"`
	FirstName  string   `json:"fname"`
	IsEnrolled bool     `json:"enrolled"`
	Courses    []Course `json:"classes"`
}

type Course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

var s Student

func simpleDecode(data []byte) {

	if !json.Valid(data) {
		log.Fatalln("The JSON data is not valid", string(data))
	}

	json.Unmarshal(data, &s)

	fmt.Println(s)

}
