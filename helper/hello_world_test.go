package helper

import (
	"fmt"
	"testing"
)

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
