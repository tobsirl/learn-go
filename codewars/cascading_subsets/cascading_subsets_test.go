package cascadingsubsets

import (
	"reflect"
	"testing"
)

func TestEachCons(t *testing.T) {
	tests := []struct {
		arr      []int
		n        int
		expected [][]int
	}{
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{[]int{1, 2, 3, 4}, 3, [][]int{{1, 2, 3}, {2, 3, 4}}},
		{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
	}

	for _, test := range tests {
		result := EachCons(test.arr, test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("EachCons(%v, %d) = %v; expected %v", test.arr, test.n, result, test.expected)
		}
	}
}	