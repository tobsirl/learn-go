package removestringspaces

import "testing"

func TestNoSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloworld"},
		{"  leading and trailing spaces  ", "leadingandtrailingspaces"},
		{"no spaces", "nospaces"},
		{"   ", ""},
		{"", ""},
	}

	for _, test := range tests {
		result := NoSpace(test.input)
		if result != test.expected {
			t.Errorf("NoSpace(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}