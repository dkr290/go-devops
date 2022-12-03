package hydrachat

import (
	"fmt"
	"io"
	"sync"

	"github.com/dkr290/go-devops/hydra/hlogger"
)

type room struct {
	name    string
	Msgch   chan string
	clients map[chan<- string]struct{}
	Quit    chan struct{}
	*sync.RWMutex
}

var logger = hlogger.GetInstance()

func CreateRoom(name string) *room {

	r := &room{
		name:    name,
		Msgch:   make(chan string),
		clients: map[chan<- string]struct{}{},
		Quit:    make(chan struct{}),
		RWMutex: new(sync.RWMutex),
	}

	r.Run()
	return r
}

func (r *room) AddClient(c io.ReadWriteCloser) {
	r.Lock()
	wc, done := StartClient(r.Msgch, c, r.Quit)
	r.clients[wc] = struct{}{}
	r.Unlock()

	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

func (r *room) Run() {

	logger.Println("Starting the chat room", r.name)
	go func() {
		for msg := range r.Msgch {
			r.broadcastMsg(msg)
		}
	}()

}

func (r *room) broadcastMsg(msg string) {
	r.RLock()
	fmt.Println("Received message: ", msg)
	for wc, _ := range r.clients {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
	defer r.RUnlock()
}
func (r *room) ClCount() int {
	return len(r.clients)
}

func (r *room) RemoveClient(wc chan<- string) {
	logger.Println("Removing client ")
	r.Lock()
	close(wc)
	delete(r.clients, wc)
	r.Unlock()
	select {
	case <-r.Quit:
		if len(r.clients) == 0 {
			close(r.Msgch)
		}
	default:
	}
}
