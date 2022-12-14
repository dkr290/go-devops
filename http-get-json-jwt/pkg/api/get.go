package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
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

type WordsPage struct {
	Page
	Words
}

func (w Words) GetResponse() string {
	return strings.Join(w.Words, ",")
}

func (o Occurrence) GetResponse() string {
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return strings.Join(out, ",")

}

type Response interface {
	GetResponse() string
}

// make own http get client
func (a Api) DoGetRequest(requestURL string) (Response, RequestError) {

	if _, err := url.ParseRequestURI(requestURL); err != nil {
		return nil, RequestError{

			Err: fmt.Errorf("vaidation error, url is not valid: %s", err),
		}
	}

	response, err := a.Client.Get(requestURL)

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
			Err:      fmt.Errorf("invalid output (HTTP Code %d): %s", response.StatusCode, string(body)),
		}

	}

	if !json.Valid(body) {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Errorf("io valid JSON returned"),
		}

	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Errorf("page Unmarshal error %s", err),
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
