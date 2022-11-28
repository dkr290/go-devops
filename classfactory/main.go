package main

import (
	"fmt"
	"log"

	"github.com/dkr290/go-devops/classfactory/appliances"
)

func main() {

	fmt.Println("Enter preferred appliance type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	// user input

	var myType int
	fmt.Scan(&myType)

	myAppliance, err := appliances.CreateAppliance(myType)

	if err != nil {
		log.Fatalln(err)
	} else {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	}

}
