package main

import (
	"log"
	"os"

	_ "github.com/dkr290/go-devops/rssseach/matches"
	"github.com/dkr290/go-devops/rssseach/search"
)

// init it called before main

func init() {
	log.SetOutput(os.Stdout)
}

func main() {

	//search for specific teram

	search.Run("president")
}
