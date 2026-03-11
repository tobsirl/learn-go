package centuryfromyear

import "testing"

func TestCentury(t *testing.T) {
	tests := []struct {
		year     int
		expected int
	}{
		{1705, 18},
		{1900, 19},
		{1601, 17},
		{2000, 20},
	}

	for _, test := range tests {
		result := Century(test.year)
		if result != test.expected {
			t.Errorf("Century(%d) = %d; expected %d", test.year, result, test.expected)
		}
	}
}
