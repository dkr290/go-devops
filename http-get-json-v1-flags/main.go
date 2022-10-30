package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("%s", strings.Join(w.Words, ","))
}

func (o Occurrence) GetResponse() string {
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("%s", strings.Join(out, ","))

}

type Response interface {
	GetResponse() string
}

var (
	requestURL string
	password   string
	parsedURL  *url.URL
	err        error
)

func main() {

	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&password, "password", "", "use a password to access our api")
	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Validation error: URL is not valid: %s", err)
		flag.Usage()
		os.Exit(1)
	}

	res, err := doRequest(parsedURL.String())

	if err.Err != nil {

		fmt.Printf("Error: %s (HTTP code: %d, Body:%s\n", err.Err, err.HTTPCode, err.Body)
		os.Exit(1)
	}

	if res == nil {

		log.Fatalln("No response")
	}

	fmt.Printf("Response: %s", res.GetResponse())
}

func doRequest(requestURL string) (Response, RequestError) {

	if _, err := url.ParseRequestURI(requestURL); err != nil {
		return nil, RequestError{

			Err: fmt.Errorf("Vaidation error, url is not valid: %s", err),
		}
	}

	response, err := http.Get(requestURL)

	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Err:      fmt.Errorf("http Get error: %s", err),
		}

	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Errorf("ReadAll error: %s", err),
		}
	}

	if response.StatusCode != 200 {

		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Errorf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, string(body)),
		}

	}

	if !json.Valid(body) {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Errorf("No valid JSON returned"),
		}

	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Errorf("Page Unmarshal error %s", err),
		}
	}

	switch page.Name {
	case "words":
		var words Words
		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Errorf("Words unmarshal error %s", err),
			}
		}
		return words, RequestError{Err: nil}

	case "occurrence":
		var oc Occurrence
		err = json.Unmarshal(body, &oc)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Errorf("Occurrence unmarshal error %s", err),
			}
		}

		return oc, RequestError{Err: nil}

	}

	return nil, RequestError{Err: nil}

}
