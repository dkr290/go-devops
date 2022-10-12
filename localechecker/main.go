package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	check := checkArgs(os.Args)

	okLocales := []string{"en_US", "fr_CN", "en_CN", "fr_FR", "ru_RU"}

	var found bool
	for i := 0; i < len(okLocales); i++ {
		if okLocales[i] == check {
			found = true
		}

	}

	if found {
		fmt.Println("the locale is supported")

	} else {
		fmt.Println("the locale is not supported")
	}

}

func checkArgs(args []string) string {

	for i := 0; i < len(args); i++ {
		if len(args) != 2 {
			log.Fatalln("you did not pass one argument")
		}
	}

	return args[1]
}
