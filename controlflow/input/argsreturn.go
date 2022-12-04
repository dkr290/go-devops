package main

import (
	"log"
	"os"
)

func returnArgs() []string {
	args := os.Args
	if len(args) == 1 {
		log.Fatalln("Need more argments")

	}

	return args
}
