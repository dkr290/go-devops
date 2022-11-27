package main

import (
	"net/http"

	"github.com/dkr290/go-devops/go-hydraweb/handlers"
)

var repo = handlers.NewRepo()

func main() {

	http.HandleFunc("/", repo.HandleRoot)
	http.ListenAndServe("127.0.0.1:8080", nil)

}
