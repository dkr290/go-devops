package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Please put here your age: ")

	reader := bufio.NewReader(os.Stdin)

	userinput, err := reader.ReadString('\n')

	if err != nil {
		log.Println("Error reading user input ", err)
	}

	userinput = strings.TrimSpace(userinput)
	userinput = strings.Trim(userinput, "\n")

	age, err := strconv.Atoi(userinput)

	if err != nil {
		log.Println("Error converting to int ", err)
	}

	checkAge(age)

}

func checkAge(age int) {

	if age < 5 {
		fmt.Println("Too young for the school")

	} else if age == 5 {
		fmt.Println("Go to kindergarten")
	} else if (age > 5) && (age <= 17) {
		grade := age - 5
		fmt.Printf("Go to grade %d", grade)
	} else {
		fmt.Println("Go to college")
	}
}
