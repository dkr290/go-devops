package main

import "fmt"

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {

	return x + y
}

func main() {

	fmt.Println("3 +1 = ", getSumGen(3, 1))
	fmt.Printf("3.4 + 1.2 = %f\n", getSumGen(3.4, 1.2))

}
