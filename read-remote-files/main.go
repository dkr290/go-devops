package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	err := remoteReader()
	if err != nil {
		log.Fatalln(err)
	}

}

func remoteReader() error {

	//example of remote file reading

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.internic.net/domain/named.root", nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	f, err := os.OpenFile("./names.txt", flags, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = io.Copy(f, resp.Body); err != nil {
		return errors.New("error copy to a file file")
	}

	return nil
}
