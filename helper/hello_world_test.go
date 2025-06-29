package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Harun")

	if result != "Hello Harun" {
		t.Error("result must be Hello Harun")
	}

	fmt.Println("TestHelloWorldHarun DONE")
}

func TestHelloWorldRasyid(t *testing.T) {
	result := HelloWorld("Rasyid")

	if result != "Hello Rasyid" {
		t.Fatal("result must be Hello Rasyid")
	}

	fmt.Println("TestHelloWorldRasyid DONE")
}
