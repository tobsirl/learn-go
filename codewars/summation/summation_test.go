package summation

import "testing"

func TestSummation(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{2, 3},
		{8, 36},
	}

	for _, test := range tests {
		result := Summation(test.n)
		if result != test.expected {
			t.Errorf("Summation(%d) = %d; want %d", test.n, result, test.expected)
		}
	}
}