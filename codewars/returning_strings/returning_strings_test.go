package returningstrings

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Alice", "Hello, Alice how are you doing today?"},
		{"Bob", "Hello, Bob how are you doing today?"},
		{"", "Hello,  how are you doing today?"},
	}

	for _, test := range tests {
		result := Greet(test.name)
		if result != test.expected {
			t.Errorf("Greet(%q) = %q; expected %q", test.name, result, test.expected)
		}
	}
}