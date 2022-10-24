package main

import "fmt"

type cat struct {
}

type Speaker interface {
	Speak() string
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func chatter(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {

	c := cat{}

	chatter(c)

}
