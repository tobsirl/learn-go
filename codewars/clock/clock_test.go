package clock

import "testing"

func TestPast(t *testing.T) {
	tests := []struct {
		h        int
		m        int
		s        int
		expected int
	}{
		{0, 0, 0, 0},
		{1, 0, 0, 3600000},
		{0, 1, 0, 60000},
		{0, 0, 1, 1000},
		{1, 1, 1, 3661000},
	}

	for _, test := range tests {
		result := Past(test.h, test.m, test.s)
		if result != test.expected {
			t.Errorf("Past(%d, %d, %d) = %d; expected %d", test.h, test.m, test.s, result, test.expected)
		}
	}
}