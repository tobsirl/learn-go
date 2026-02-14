package smallestintegerfinder

import "testing"

func TestFindSmallestInteger(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Test Case 1", []int{1, 2, 3, 4, 5}, 1},
		{"Test Case 2", []int{-1, -2, -3, -4, -5}, -5},
		{"Test Case 3", []int{0, 0, 0}, 0},
		{"Test Case 4", []int{10, 20, 5, 15}, 5},
		{"Test Case 5", []int{-10, -20, -5, -15}, -20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SmallestIntegerFinder(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, result)
			}
		})
	}
}