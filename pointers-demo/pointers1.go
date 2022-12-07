package main

import (
	"fmt"
	"reflect"
)

func dType(t any) {

	switch v := t.(type) {
	case string:
		t2 := v + "..."
		fmt.Println("String found", t2)
	case *string:
		fmt.Printf("Pointer string found %s\n", *v)
	case int:
		fmt.Printf("Integer foud %d", v)

	default:
		fmt.Printf("Type was not found %v\n", reflect.TypeOf(t))
	}

}
