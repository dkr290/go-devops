package main

import (
	"time"
)

func locTime(l, forparsing, input string) (time.Time, error) {
	loc, _ := time.LoadLocation(l)
	now, err := time.Parse(forparsing, input)
	if err != nil {
		return time.Now(), err
	}

	return now.In(loc), nil

}
