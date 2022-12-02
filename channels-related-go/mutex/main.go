package main

import (
	"fmt"
	"sync"
)

type mytype struct {
	counter int
	mu      sync.Mutex
}

func main() {
	mtInstance := mytype{}
	finished := make(chan bool)

	for i := 0; i < 5; i++ {
		go func(m *mytype) {
			mtInstance.mu.Lock()

			fmt.Printf("input counter %d\n", m.counter)
			m.counter++
			fmt.Printf("output counter %d\n", m.counter)
			finished <- true
			mtInstance.mu.Unlock()

		}(&mtInstance)

	}
	for i := 0; i < 5; i++ {
		<-finished
	}

	fmt.Println("Counter", mtInstance.counter)

}
