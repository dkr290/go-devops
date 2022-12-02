package main

import (
	"fmt"
	"time"
)

func main() {

	go slowCounter(2)
	time.Sleep(15 * time.Second)

}

func slowCounter(n int) {
	i := 0
	//create a duration if n seconds
	d := time.Duration(n) * time.Second

	for {
		//create a timer for this duration
		t := time.NewTimer(d)
		<-t.C
		i++
		fmt.Println(i)
	}
}
