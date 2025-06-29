package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Harun")

	if result != "Hello Harun" {
		panic("result not Hello Harun")
	}
}
