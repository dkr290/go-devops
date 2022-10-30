package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/dkr290/go-devops/http-get-json-jwt/pkg/api"
)

var (
	requestURL string
	password   string
	parsedURL  *url.URL
	err        error
)

func main() {

	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&password, "p", "", "use a password to access our api")
	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Validation error: URL is not valid: %s", err)
		flag.Usage()
		os.Exit(1)
	}

	apiInstance := api.New(api.Options{
		Password: password,
		LoginUrl: parsedURL.Scheme + "://" + parsedURL.Host + "/login",
	})

	res, err := apiInstance.DoGetRequest(parsedURL.String())

	if err.Err != nil {

		fmt.Printf("Error: %s (HTTP code: %d, Body:%s\n", err.Err, err.HTTPCode, err.Body)
		os.Exit(1)
	}

	if res == nil {

		log.Fatalln("No response")
	}

	fmt.Printf("Response: %s", res.GetResponse())
}
