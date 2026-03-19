package reducebutgrow

import "testing"

func TestGrow(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{4, 5, 6}, 120},
		{[]int{7, 8, 9}, 504},
		{[]int{}, 1},
	}

	for _, test := range tests {
		result := Grow(test.arr)
		if result != test.expected {
			t.Errorf("Grow(%v) = %v; expected %v", test.arr, result, test.expected)
		}
	}
}