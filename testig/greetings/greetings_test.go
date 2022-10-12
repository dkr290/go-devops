package greetings

import (
	"fmt"
	"testing"
)

func Greet(name string) (string, error) {

	if name == "" {
		return "", fmt.Errorf("name was empty")
	} else {
		return "Hello " + name, nil
	}

}

func TestGreet(t *testing.T) {
	name := "Bob"
	want := "Hello Bob"

	got, err := Greet(name)
	if got != want || err != nil {
		t.Fatalf("TestGreet(%s): got %q/%v, want %q/nil", name, got, err, want)
	}

}
