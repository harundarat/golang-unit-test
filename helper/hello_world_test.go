package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
