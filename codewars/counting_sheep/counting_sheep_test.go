package countingsheep

import "testing"

func TestCountSheep(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []bool
		expected int
	}{
		{"Example 1", []bool{true, true, true, false}, 3},
		{"Example 2", []bool{true, false, true, false, true}, 3},
		{"Example 3", []bool{false, false, false}, 0},
		{"Example 4", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountSheep(tt.numbers)
			if result != tt.expected {
				t.Errorf("Expected %d but got %d", tt.expected, result)
			}
		})
	}
}