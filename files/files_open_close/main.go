package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	iPrimeArr := []int{1, 2, 3, 4, 7, 11, 34}
	var sArr []string

	for _, val := range iPrimeArr {
		sArr = append(sArr, strconv.Itoa(val))
	}

	for _, num := range sArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan1 := bufio.NewScanner(f)

	for scan1.Scan() {
		fmt.Println("Prime : ", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exeist")
	} else {
		f, err = os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("18\n"); err != nil {
			log.Fatal(err)
		}
	}

}
