package returnnegative

import "testing"

func TestMakeNegative(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, -1},
		{-5, -5},
		{0, 0},
	}

	for _, test := range tests {
		result := MakeNegative(test.input)
		if result != test.expected {
			t.Errorf("MakeNegative(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}