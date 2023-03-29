package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldJoko(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Joko")
	}
}

func BenchmarkHelloWorldSusilo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Susilo")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
		{
			name:     "Susilo",
			request:  "Susilo",
			expected: "Hello Susilo",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Joko", func(t *testing.T) {
		result := HelloWorld("Joko")
		require.Equal(t, "Hello Joko", result, "Result must be 'Hello Joko'")
	})

	t.Run("Susilo", func(t *testing.T) {
		result := HelloWorld("Susilo")
		require.Equal(t, "Hello Susilo", result, "Result must be 'Hello Susilo'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Joko")
	require.Equal(t, "Hello Joko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Joko")
	assert.Equal(t, "Hello Joko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldJoko(t *testing.T) {
	result := HelloWorld("Joko")

	if result != "Hello Joko" {
		// failed test
		t.Error("Result must be 'Hello Joko'")
	}

	fmt.Println("TestHelloWorldJoko DONE")
}

func TestHelloWorldSusilo(t *testing.T) {
	result := HelloWorld("Susilo")

	if result != "Hello Susilo" {
		// failed test
		t.FailNow()
	}
}
