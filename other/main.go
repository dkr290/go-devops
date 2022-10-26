package main

import "fmt"

type cat struct {
	name string
}

func main() {

	c := cat{name: "some name"}
	i := []interface{}{42, "The book club", true, c}
	typeExample(i)

}

func typeExample(i []interface{}) {

	for _, val := range i {

		switch v := val.(type) {
		case int:
			fmt.Printf("%v is of type int\n", v)
		case string:
			fmt.Printf("%v is of type string\n", v)

		case bool:
			fmt.Printf("%v is of type bool\n", v)
		default:
			fmt.Printf("Unknown type\n")
		}

	}

}
