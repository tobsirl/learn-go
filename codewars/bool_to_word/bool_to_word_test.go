package booltoword

import "testing"

func TestBoolToWord(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{true, "Yes"},
		{false, "No"},
	}

	for _, test := range tests {
		result := BoolToWord(test.input)
		if result != test.expected {
			t.Errorf("BoolToWord(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}