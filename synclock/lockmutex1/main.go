package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	i int
	sync.Mutex
}

func (sc *safeCounter) Increment() {
	sc.Mutex.Lock()
	sc.i++
	sc.Mutex.Unlock()

}

func (sc *safeCounter) Decrement() {
	sc.Lock()
	sc.i--
	sc.Unlock()
}

func main() {
	sc := new(safeCounter)

	for i := 0; i < 100; i++ {
		go sc.Increment()
		go sc.Decrement()
	}

	fmt.Println(sc.GetValue())

}

func (sc *safeCounter) GetValue() int {
	sc.Mutex.Lock()
	v := sc.i
	sc.Mutex.Unlock()
	return v

}
