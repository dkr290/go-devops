package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type getReq struct {
	Author string `json:"author"`
}

type getResp struct {
	Quote string `json:"quote"`
	Error *Error `json:"error"`
}

type Error struct {
	Code string
	Msg  string
}

type QOTD struct {
	addr   string
	client *http.Client
}

func (e *Error) Error() error {
	return fmt.Errorf("(code %v): %s", e.Code, e.Msg)
}

func New(addr string) (*QOTD, error) {
	if _, _, err := net.SplitHostPort(addr); err != nil {
		return nil, err
	}
	return &QOTD{
		addr:   addr,
		client: &http.Client{},
	}, nil
}

func main() {

}

func (q *QOTD) restCall(ctx context.Context, endpoint string, req, resp interface{}) error {

}
