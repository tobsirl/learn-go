package helloworld

import "testing"

func TestGreet(t *testing.T) {
	expected := "hello, world!"
	result := Greet()
	if result != expected {
		t.Errorf("Greet() = %v; want %v", result, expected)
	}
}