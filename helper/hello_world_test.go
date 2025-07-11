package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{name: "Harun", request: "Harun"},
		{name: "Al", request: "Al"},
		{name: "Rasyid", request: "Rasyid"},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Harun", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Harun")
		}
	})
	b.Run("Al", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Al")
		}
	})
	b.Run("Rasyid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rasyid")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Harun")
	}
}
func BenchmarkHelloWorldRasyid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rasyid")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{name: "Harun", request: "Harun", expected: "Hello Harun"},
		{name: "Rasyid", request: "Rasyid", expected: "Hello Rasyid"},
		{name: "Al", request: "Al", expected: "Hello Al"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

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
