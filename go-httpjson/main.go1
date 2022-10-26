package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

/*
{
	"id": 2,
	"title": "Something more to do today",
	"completed": false
},
*/

type Todos struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	args := os.Args

	if len(args) < 2 {
		log.Fatalf("Usage: ./http-get <url>\n")

	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		log.Fatalf("URL is in invalid format: %s\n", err)
	}

	res, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {

		log.Fatalf("Invalid output (HTTP code %d): %s\n", res.StatusCode, string(body))
	}

	var todos []Todos

	err = json.Unmarshal(body, &todos)
	if err != nil {
		log.Fatal("Error Unmarshal ", err)
	}

	for _, todo := range todos {

		fmt.Printf("JSON parsed Title: %s , Completed: %v\n", todo.Title, todo.Completed)

	}
}
