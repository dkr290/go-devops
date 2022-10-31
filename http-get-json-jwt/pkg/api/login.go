package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func doLoginRequest(client http.Client, requestURL, password string) (string, RequestError) {

	loginRequest := LoginRequest{
		Password: password,
	}

	b, err := json.Marshal(loginRequest)
	if err != nil {
		return "", RequestError{Err: fmt.Errorf("marshal error: %s", err)}
	}

	bf := bytes.NewBuffer(b)
	response, err := client.Post(requestURL, "application/json", bf)
	if err != nil {
		return "", RequestError{
			HTTPCode: response.StatusCode,
			Err:      fmt.Errorf("http Post error: %s", err),
		}

	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return "", RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(responseBody),
			Err:      fmt.Errorf("ReadAll error: %s", err),
		}
	}

	if response.StatusCode != 200 {

		return "", RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(responseBody),
			Err:      fmt.Errorf("invalid output (HTTP Code %d): %s", response.StatusCode, string(responseBody)),
		}

	}

	if !json.Valid(responseBody) {
		return "", RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(responseBody),
			Err:      fmt.Errorf("no valid JSON returned"),
		}

	}

	var lResponse LoginResponse
	err = json.Unmarshal(responseBody, &lResponse)
	if err != nil {
		return "", RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(responseBody),
			Err:      fmt.Errorf("Page Unmarshal error %s", err),
		}
	}

	if lResponse.Token == "" {
		return "", RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(responseBody),
			Err:      fmt.Errorf("empty token replied"),
		}
	}

	return lResponse.Token, RequestError{Err: nil}
}
