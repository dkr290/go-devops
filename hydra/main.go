package main

import (
	"fmt"
	"net/http"

	"github.com/dkr290/go-devops/hydra/hlogger"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	http.HandleFunc("/", handleRoot)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprintf(w, "welcome to the Hydra software system")

	logger.Println("Received an http Get request on root url")
}
