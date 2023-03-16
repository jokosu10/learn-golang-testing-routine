package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Joko")

	if result != "Hello Joko" {
		// failed test
		panic("Result is no 'Hello Eko'")
	}
}
