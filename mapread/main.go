package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("User id not passed")
	}

	userid := os.Args[1]
	name, err := getUser(userid)
	if err != nil {
		fmt.Println("The user does not match in the database")
		for k, v := range getUsers() {
			fmt.Println("id:  :", k, "Name:  ", v)
		}
		log.Fatalf(err.Error())
	}

	fmt.Println("Name:  ", name)

}

func getUsers() map[string]string {

	return map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
}

func getUser(id string) (string, error) {
	users := getUsers()
	if user, ok := users[id]; ok {
		return user, nil
	} else {
		return "", errors.New("There is and error the user id cannot be found ")
	}
}
