package convertnumbertoreversedarrayofdigits

import "testing"

func TestDigitize(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{input: 35231, expected: []int{1, 3, 2, 5, 3}},
		{input: 0, expected: []int{0}},
	}

	for _, test := range tests {
		result := Digitize(test.input)
		if !equal(result, test.expected) {
			t.Errorf("Digitize(%d) = %v; expected %v", test.input, result, test.expected)
		}
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