package main

import (
	"errors"
	"fmt"
)

func main() {

	res, _ := doubleR(5)
	fmt.Println(res)
	res, _ = doubleR(36.77676)
	fmt.Println(res)
	res, _ = doubleR("Hello Tom, how are you")
	fmt.Println(res)

}

func doubleR(v interface{}) (string, error) {

	switch t := v.(type) {
	case int32, int64, int:
		if n, ok := t.(int32); ok {
			return fmt.Sprintf("The type is %T", n), nil
		}
		if n, ok := t.(int64); ok {
			return fmt.Sprintf("The type is %T", n), nil
		}
		if n, ok := t.(int); ok {
			return fmt.Sprintf("The type is %T", n), nil
		}

	case string:
		return fmt.Sprintf("The type is %T", t), nil

	case bool:
		return fmt.Sprintf("The type is %T", t), nil

	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprintf("The type is %T", f), nil
		}
		if f, ok := t.(float32); ok {
			return fmt.Sprintf("The type is %T", f), nil
		}
	default:
		return "", errors.New("There is unsupported type")
	}
	return "", nil
}
