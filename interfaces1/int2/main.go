package main

import (
	"fmt"
	"math/rand"
	"time"
)

type custRand struct {
	*rand.Rand
	count int
}

func newCustRand(i int64) *custRand {

	return &custRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *custRand) RandRange(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

func (cr *custRand) getCount() int {
	return cr.count
}

func main() {

	cr := newCustRand(time.Now().UnixMicro())
	fmt.Println(cr.RandRange(5, 30))

	fmt.Println(cr.Intn(10))
	fmt.Println(cr.getCount())
}
