package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	ResponseOutput *http.Response
}

func (m *MockClient) Get(url string) (resp *http.Response, err error) {

	return m.ResponseOutput, nil

}

func TestDoGetRequest(t *testing.T) {
	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}
	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Errorf("marshal error: %s", err)
	}

	apiInstance := Api{
		Options: Options{},
		Client: &MockClient{
			ResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}

	response, req_err := apiInstance.DoGetRequest("http://localhost/wordss")
	if req_err.Err != nil {
		t.Errorf("DoRequest error %s,", req_err.Err)
	}

	if response == nil {
		t.Fatalf("response is empty")
	}

	if response.GetResponse() != strings.Join([]string{"a", "b"}, ",")
}
