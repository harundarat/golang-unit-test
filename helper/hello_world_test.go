package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Harun", func(t *testing.T) {
		result := HelloWorld("Harun")
		assert.Equal(t, "Hello Harun", result, "Result must be 'Hello Harun'")
	})

	t.Run("Rasyid", func(t *testing.T) {
		result := HelloWorld("Rasyid")
		assert.Equal(t, "Hello Rasyid", result, "Result must be 'Hello Rasyid'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on macos")
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

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Harun")
	require.Equal(t, "Hello Harun", result, "Result must be 'Hello Harun'")
	fmt.Println("This won't be executed")
}
