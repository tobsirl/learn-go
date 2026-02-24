package youcantcodeunderpressure

import "testing"

func TestDoubleInteger(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 4},
		{0, 0},
		{-1, -2},
	}

	for _, test := range tests {
		result := DoubleInteger(test.input)
		if result != test.expected {
			t.Errorf("DoubleInteger(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}