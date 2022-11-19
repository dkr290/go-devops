package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CreateFile(fname string) {

	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("Using write function\n"))
	f.WriteString("Using writesting function\n")
}

func Wfilw(fname string) {

	m := []byte("This is another test string to file to add inside")

	err := ioutil.WriteFile(fname, m, 0644)
	if err != nil {
		log.Fatalln("An error creating the file")
	}

}

func RAll(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File Contents\n", string(content))

	r := strings.NewReader("No file from this file new")
	c, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Content of the strings new reader is", string(c))

	f1, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}
	defer f1.Close()

	if _, err := f1.Write([]byte("Writing some stuff")); err != nil {
		panic(err)
	}
}
