package main

import "fmt"

type Store struct {
	value interface{}
	name  string
}

func (s *Store) setStore(v interface{}) {
	s.value = v

}

func (s *Store) getStore() any {

	return s.value

}

func NewStore(v string) *Store {

	return &Store{name: v}
}

func main() {

	intstore := NewStore("Integers to store")
	intstore.setStore(4)
	if v, ok := intstore.getStore().(int); ok {
		v *= 10
		fmt.Println(v)
	}

	stringstore := NewStore("Strings to store")
	stringstore.setStore("Hello text")
	if v, ok := stringstore.getStore().(string); ok {
		v += " Some other text"
		fmt.Println(v)
	}
}
