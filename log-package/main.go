package main

import (
	"github.com/dkr290/go-devops/log-package/logging"
)

func main() {

	logging.Trace.Println("I have something standart to say")
	logging.Info.Println("Special information")
	logging.Warning.Println("There is something you need to know about")
	logging.Error.Println("Somethibng has failed")

}
