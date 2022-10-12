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
	tests := []struct {
		description string
		name        string
		want        string
		expectErr   bool
	}{
		{
			description: "Error: name is an empty string",
			expectErr:   true,
			//name and want "", the zero value for string
		},
		{
			description: "Sucess",
			name:        "John",
			want:        "Hello John",
			//Expected error is false
		},
	}

	for _, test := range tests {
		got, err := Greet(test.name)

		switch {
		// dont get and error but expect one
		case err == nil && test.expectErr:
			t.Errorf("TestGreet %s: got err == nil, want err != nil", test.description)

			// got un error but do not expect one
		case err != nil && !test.expectErr:
			t.Errorf("TestGreet %s: got err == %s , want err == nil", test.description, err)

		case err != nil:
			continue

		}
		if got != test.want {
			t.Errorf("TestGret %s: got result %v, want %v", test.description, got, test.want)
		}

	}

}
