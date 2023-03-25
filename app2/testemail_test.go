package app2

import "testing"

func TestIsEmail(t *testing.T) {

	_, err := IsEmail("hello@gmail.com")

	if err != nil {
		t.Error("hello@gmail.com is an email")
	}

	_, err = IsEmail("hello@gmail")

	if err != nil {
		t.Error("hello@gmail is an email")
	}

}
