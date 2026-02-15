package stringtonumber

import "testing"

func TestStringToNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1234", 1234},
		{"5678", 5678},
		{"0", 0},
		{"42", 42},
		{"-7", -7},
	}

	for _, test := range tests {
		result := StringToNumber(test.input)
		if result != test.expected {
			t.Errorf("StringToNumber(%q) = %d; expected %d", test.input, result, test.expected)
		}
	}
}