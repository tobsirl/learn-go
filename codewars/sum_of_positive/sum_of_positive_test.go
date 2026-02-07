package sumofpositive

import "testing"

func TestPositiveSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, -4, 7, 12}, 20},
		{[]int{-1, -2, -3, -4, -5}, 0},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := PositiveSum(test.input)
		if result != test.expected {
			t.Errorf("PositiveSum(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}