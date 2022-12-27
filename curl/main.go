package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dkr290/go-devops/curl/parsing"
)

var r, f string

func init() {
	flag.StringVar(&r, "r", "", "go-curl -r url -f file")
	flag.StringVar(&f, "f", "", "go-curl -r url -f file")
}
func main() {

	flag.Parse()

	if len(os.Args) != 5 {

		flag.PrintDefaults()
	}

	res, err := http.Get(r)
	if err != nil {
		log.Fatalln(err)
	}

	dest, file := parsing.CustomWrite(f)

	defer file.Close()

	io.Copy(dest, res.Body)

	if err := res.Body.Close(); err != nil {
		log.Println(err)
	}

}
