package main

import (
	"fmt"
	"log"
	"time"
)

func timeDiff(timezone string) (string, string) {

	Current := time.Now()

	RemoteZone, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(RemoteZone)
	RemoteTime := Current.In(RemoteZone)

	fmt.Println("The current time is:", Current.Format(time.ANSIC))
	fmt.Println("The timezone:", timezone, "time is", RemoteTime)

	return Current.Format(time.ANSIC), RemoteTime.Format(time.ANSIC)

}
