package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on linux")
	}

	fmt.Println("runtime: " + runtime.GOOS)
	result := HelloWorld("Harun")
	assert.Equal(t, "Hello Harun", result, "Result must be 'Hello Harun'")
	fmt.Println("This will be executed")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Harun")
	assert.Equal(t, "Hello Harun", result, "Result must be 'Hello Harun'")
	fmt.Println("This will be executed")
}

func TestHElloWorldRequire(t *testing.T) {
	result := HelloWorld("Harun")
	require.Equal(t, "Hi Harun", result, "Result must be 'Hello Harun'")
	fmt.Println("This won't be executed")
}
