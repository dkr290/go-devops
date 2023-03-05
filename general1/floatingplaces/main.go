package main

import "fmt"

func main() {

	fmt.Printf("%9f\n", 3.18)
	fmt.Printf("%.3f\n", 3.3224343)
	fmt.Printf("%9.f\n", 3.14675675)
	sp1 := fmt.Sprintf("%9.f\n", 3.1454645)
	fmt.Println(sp1)
}
