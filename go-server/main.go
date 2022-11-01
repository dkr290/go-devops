package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request sucessful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name %s\n", name)
	fmt.Fprintf(w, "Address %s\n", address)
}

func envHandler(w http.ResponseWriter, r *http.Request) {

	dbh := os.Getenv("DATABASE_HOST")
	dbp := os.Getenv("DATABASE_PASS")
	dbuser := os.Getenv("DATABASE_USER")

	if r.URL.Path != "/env" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "methgod is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "DATABASE USER: %s\n DATABASE PASSWORD: %s\n DATABASE HOST: %s\n", dbuser, dbp, dbh)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/env", envHandler)

	fmt.Printf("Starting the server at port 8080\n")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln(err)
	}

}
