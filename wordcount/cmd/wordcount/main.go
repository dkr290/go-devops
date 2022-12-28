package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dkr290/go-devops/wordcount"
)

func main() {

	var file string
	flag.StringVar(&file, "f", "", "worldcount -f <filename>")

	flag.Parse()

	if len(os.Args) != 3 {
		flag.Usage()
		return
	}
	wordCount, err := wordcount.CountWordsInFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Found %d wordcount", wordCount)

}
