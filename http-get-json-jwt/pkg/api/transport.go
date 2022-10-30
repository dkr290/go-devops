package api

import (
	"net/http"
)

type MyJWTTransport struct {
	transport http.RoundTripper
	token     string
	password  string
	loginURL  string
}

// calling the function but overwriting default behavious adding header
func (m *MyJWTTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if m.token == "" {
		if m.password != "" {
			token, err := doLoginRequest(http.Client{}, m.loginURL, m.password)
			if err.Err != nil {
				return nil, err.Err

			}

			m.token = token
		}
	}
	if m.token != "" {
		r.Header.Add("Authorization", "Bearer "+m.token)
	}

	return m.transport.RoundTrip(r)
}
