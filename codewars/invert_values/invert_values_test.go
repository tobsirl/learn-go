package invertvalues

import "testing"

func TestInvert(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"Example test case", []int{1, 2, 3, 4, 5}, []int{-1, -2, -3, -4, -5}},
		{"Test with negative numbers", []int{-1, -2, -3}, []int{1, 2, 3}},
		{"Test with mixed numbers", []int{-1, 0, 1}, []int{1, 0, -1}},
		{"Test with empty array", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Invert(tt.arr); !equal(got, tt.want) {
				t.Errorf("Invert(%v) = %v; want %v", tt.arr, got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}