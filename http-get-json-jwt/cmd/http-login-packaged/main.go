package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
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

	client := http.Client{}

	if password != "" {
		token, err := api.doLoginRequest(client, parsedURL.Scheme+"://"+parsedURL.Host+"/login", password)
		if err.Err != nil {

			fmt.Printf("Error: %s (HTTP code: %d, Body:%s\n", err.Err, err.HTTPCode, err.Body)
			os.Exit(1)
		}

		client.Transport = &pkg.api.myJWTTransport{
			transport: http.DefaultTransport,
			token:     token,
		}
	}

	res, err := doRequest(client, parsedURL.String())

	if err.Err != nil {

		fmt.Printf("Error: %s (HTTP code: %d, Body:%s\n", err.Err, err.HTTPCode, err.Body)
		os.Exit(1)
	}

	if res == nil {

		log.Fatalln("No response")
	}

	fmt.Printf("Response: %s", res.GetResponse())
}
