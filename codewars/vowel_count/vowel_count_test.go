package vowelcount

import "testing"

func TestGetCount(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{"hello", 2},
		{"world", 1},
		{"golang", 2},
		{"", 0},
	}

	for _, test := range tests {
		result := GetCount(test.str)
		if result != test.expected {
			t.Errorf("GetCount(%q) = %v; expected %v", test.str, result, test.expected)
		}
	}
}