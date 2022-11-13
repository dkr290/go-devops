package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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

	s := NewStudent(2, "Miarsham", "s", "John", false, false)

	c := CourseMarshal{Name: "Mathematic", Number: 101, Hours: 3}
	s.Courses = append(s.Courses, c)
	c = CourseMarshal{Name: "Biology", Number: 234, Hours: 8}
	s.Courses = append(s.Courses, c)

	mStudent, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(mStudent))

	jsonData01 := []byte(`
	{
	   "id": 2,
	   "lname": "Washington",
	   "fname": "Johns",
	   "IsEnrolled": true,
	   "grades": [100,76,93,50],
	   "class": 
	   {
		"coursename": "Mathematics",
		"coursenum": 101,
		"coursehours": 3
	   }
	}
	`)

	if !json.Valid(jsonData01) {
		fmt.Printf("Json is not valid: %s", jsonData01)
		os.Exit(1)
	}

	var v interface{}
	err = json.Unmarshal(jsonData01, &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data1 := v.(map[string]interface{})

	for k, v := range data1 {
		switch val := v.(type) {
		case string:
			fmt.Println("(string):", k, val)
		case float64:
			fmt.Println("(float64):", k, val)
		case bool:
			fmt.Println("(bool):", k, val)

		case []interface{}:
			fmt.Println("(slice):", k)

			for i, j := range val {
				fmt.Println("    ", i, j)
			}
		case map[string]interface{}:
			for i, v := range val {
				fmt.Println("Course", i, ":", v)
			}

		default:
			fmt.Println("(unknown):", k, val)
		}
	}

}
