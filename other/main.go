package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func main() {

	var jsonParsed any
	err := json.Unmarshal([]byte(`{"test": {"test2":[1,2,3] },"test3": "..." }`), &jsonParsed)
	if err != nil {
		log.Fatal(err)
	}

	switch v := jsonParsed.(type) {
	case map[string]any:
		fmt.Printf("Map found: %v\n", v)
		field1, ok := v["test"]
		if ok {
			switch v2 := field1.(type) {
			default:
				fmt.Printf("Type not found %v\n", reflect.TypeOf(v2))

			}
		}

	default:
		fmt.Printf("Type not found %v\n", reflect.TypeOf(jsonParsed))

	}

}
