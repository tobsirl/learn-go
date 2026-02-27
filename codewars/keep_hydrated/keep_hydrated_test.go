package keephydrated

import "testing"

func TestLiters(t *testing.T) {
	tests := []struct {
		time     float64
		expected int
	}{
		{2, 1},
		{1.4, 0},
		{12.3, 6},
	}

	for _, test := range tests {
		result := Liters(test.time)
		if result != test.expected {
			t.Errorf("Liters(%f) = %d; expected %d", test.time, result, test.expected)
		}
	}
}