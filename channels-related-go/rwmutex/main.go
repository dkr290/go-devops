package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MapCounter struct {
	sync.RWMutex
	m map[int]int
}

func main() {

	mc := MapCounter{m: make(map[int]int)}
	go runWrites(&mc, 10)
	go runReads(&mc, 10)
	go runReads(&mc, 10)
	time.Sleep(10 * time.Second)

}

func runWrites(mc *MapCounter, n int) {
	for i := 0; i < n; i++ {
		mc.RWMutex.Lock()
		mc.m[i] = i * 10
		mc.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func runReads(mc *MapCounter, n int) {

	for {
		mc.RLock()
		v := mc.m[rand.Intn(n)]
		mc.RUnlock()
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}
