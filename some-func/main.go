package main

import (
	"fmt"
	"runtime"
)

func main() {

	//functions as values

	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))

	fmt.Println(computeMultiplyVal(3, add))
	// by passing add we are passing func(a,b int) which returns a+b

	i := increment()
	fmt.Println(i(), i(), i())

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows os")
		//fallthrough
	case "linux":
		fmt.Println("Linux os")

	default:
		fmt.Printf("%s\n", os)
		defer fmt.Println("Exiting the function...")
		fmt.Println("Entering the function")

	}

}

// passing function to another one as parameter

func computeMultiplyVal(val int, fn func(a, b int) int) int {
	return val * fn(val, val)

	// it fn return a+b but in this case is 3 +3 not the above from main

}

//function closures

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
