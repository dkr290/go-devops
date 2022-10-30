package api

import "net/http"

type myJWTTransport struct {
	transport http.RoundTripper
	token     string
}

//calling the function but overwriting default behavious adding header
func (m *myJWTTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if m.token != "" {
		r.Header.Add("Authorization", "Bearer "+m.token)
	}

	return m.transport.RoundTrip(r)
}
