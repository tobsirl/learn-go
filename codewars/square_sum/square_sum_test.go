package squaresum

import "testing"

func TestSquareSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{"Test Case 1", []int{1, 2, 3}, 14},
		{"Test Case 2", []int{0, 3, 4}, 25},
		{"Test Case 3", []int{-1, -2, -3}, 14},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SquareSum(test.input)
			if result != test.output {
				t.Errorf("Expected %d but got %d", test.output, result)
			}
		})
	}
}