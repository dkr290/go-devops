package main

import (
	"strconv"
	"time"
)

func elapsedTime(start time.Time, end time.Time) string {
	Elapsed := end.Sub(start)
	Hours := strconv.Itoa(int(Elapsed.Hours()))
	Minutes := strconv.Itoa(int(Elapsed.Minutes()))
	Seconds := strconv.Itoa(int(Elapsed.Seconds()))

	return "the total time duration elapsed for the function is: " + Hours + " hours " + Minutes + " minutes  and " + Seconds + " seconds"
}
