package main

import (
	"net/http"

	"github.com/dkr290/go-devops/go-hydraweb/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HandleRoot)
	http.ListenAndServe("127.0.0.1:8080", nil)

}
