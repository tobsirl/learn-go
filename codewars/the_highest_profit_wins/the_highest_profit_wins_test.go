package thehighestprofitwins

import "testing"

func TestMinMax(t *testing.T) {
	tests := []struct {
		arr      []int
		expected [2]int
	}{
		{[]int{1, 2, 3, 4, 5}, [2]int{1, 5}},
		{[]int{-1, -2, -3, -4, -5}, [2]int{-5, -1}},
		{[]int{5}, [2]int{5, 5}},
		{[]int{}, [2]int{0, 0}},
	}

	for _, test := range tests {
		result := MinMax(test.arr)
		if result != test.expected {
			t.Errorf("MinMax(%v) = %v; expected %v", test.arr, result, test.expected)
		}
	}
}