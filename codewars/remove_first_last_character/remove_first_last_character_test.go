package removefirstlastcharacter

import "testing"

func TestRemoveChar(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"eloquent", "loquen"},
		{"country", "ountr"},
		{"person", "erso"},
		{"place", "lac"},
		{"ok", ""},
	}

	for _, test := range tests {
		result := RemoveChar(test.input)
		if result != test.expected {
			t.Errorf("RemoveChar(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}