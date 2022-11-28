package main

import "fmt"

type cat struct{}
type dog struct{}
type person struct {
	name string
}

type Speaker interface {
	Speak() string
}

func main() {

	c := cat{}
	d := dog{}
	p := person{name: "James"}

	theSpeak(c, d, p)

	intops()
	list := NewSingleLinkedList()
	list.Add(8)
	list.Add(9)

	fmt.Println(list)

}

func (c cat) Speak() string {

	return "Purr Meow"
}

func (d dog) Speak() string {
	return "Woof Woof"
}

func (p person) Speak() string {
	return "Hi my name is " + p.name
}

func theSpeak(say ...Speaker) {
	for _, s := range say {
		fmt.Println(s.Speak())
	}
}
