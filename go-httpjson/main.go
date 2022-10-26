package main

import (
	"fmt"
	"io"
	"log"
)

/*
{
	"id": 2,
	"title": "Something more to do today",
	"completed": false
},
*/

type MySlowReader struct {
	Contents string
	Pos      int
}

// if it is not the pointer POS is always 0 because it is a copy when passwd the struct to the function
func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.Pos+1 <= len(m.Contents) {
		n := copy(p, m.Contents[m.Pos:m.Pos+1])
		m.Pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {

	ms := MySlowReader{
		Contents: "hello world",
	}

	out, err := io.ReadAll(&ms)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Output:", string(out))

}
