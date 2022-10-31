package api

import "net/http"

type Options struct {
	Password string
	LoginUrl string
}

type ClientInterface interface {
	Get(url string) (resp *http.Response, err error)
}

type APIInterface interface {
	DoGetRequest(requestURL string) (Response, RequestError)
}

type Api struct {
	Options Options
	Client  ClientInterface
}

func New(options Options) APIInterface {
	return Api{
		Options: options,
		Client: &http.Client{
			Transport: &MyJWTTransport{
				transport: http.DefaultTransport,
				password:  options.Password,
				loginURL:  options.LoginUrl,
			},
		},
	}

}
