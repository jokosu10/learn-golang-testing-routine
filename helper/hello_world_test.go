package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Joko")
	require.Equal(t, "Hi Joko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Joko")
	assert.Equal(t, "Hi Joko", result, "Result must be 'Hello Eko'")
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
