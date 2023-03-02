package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println("Hello there")
	fmt.Print("What is your name ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Hello", name)

	//some conversion

	cv1 := 1.5

	cv2 := int(cv1)
	fmt.Printf("THe %d and type %v\n", cv2, reflect.TypeOf(cv2))

	cv3 := "5000000"

	cv4, err := strconv.Atoi(cv3)
	if err != nil {
		log.Fatal("Error conversion string", err)
	}

	fmt.Printf("The value is %d  and type %v\n", cv4, reflect.TypeOf(cv4))

	cv7 := 500000000

	cv8 := strconv.Itoa(cv7)

	fmt.Println("The string is", cv8)

}
