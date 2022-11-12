package main

import "fmt"

func main() {
	unmarsh1()
	data := []byte(`{

		"id":123,
		"lname":"Bond",
		"minitial":"null",
		"fname":"James",
		"enrolled":true,
		"classes":[{
			"coursename":"Intro to golang",
			"coursenum":101,
			"coursehours":4
		},
		{
			"coursename":"Intro to python",
			"coursenum":108,
			"coursehours":3
		},
		{
			"coursename":"Mathematics",
			"coursenum":102,
			"coursehours":9
		}]
	}
	`)
	simpleDecode(data)
	fmt.Println("")
	marshal1()
}
