package main

import (
	"fmt"
	"log"
	"strconv"
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

func formatTime() {
	Current := time.Now()
	t := time.Date(Current.Year(), Current.Month(), Current.Day(), Current.Hour(), Current.Minute(), Current.Second(), 0, time.UTC)
	fmt.Printf("%d:%d:%d %d/%d/%d\n", t.Hour(), t.Minute(), t.Second(), t.Year(), t.Month(), t.Day())

}

func formatTime1() {

	date := time.Now()

	fmt.Println(strconv.Itoa(date.Hour()) + ":" + strconv.Itoa(date.Minute()) + ":" + strconv.Itoa(date.Second()) + " " + strconv.Itoa(date.Year()) + "/" + strconv.Itoa(date.Day()) + "/" + strconv.Itoa(int(date.Month())))
}

func timeRun() time.Duration {
	start := time.Now()
	time.Sleep(3 * time.Second)
	end := time.Now()
	diff := end.Sub(start)
	return diff
}
