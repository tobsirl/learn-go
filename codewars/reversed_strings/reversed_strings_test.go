package reversedstrings

import (
	"testing"
)

func TestReversedStrings(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"world", "dlrow"},
		{"word", "drow"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"  leading and trailing  ", "  gniliart dna gnidael  "},
		{"Hello, 世界", "界世 ,olleH"},
	}

	for _, tc := range tests {
		got := Solution(tc.in)
		if got != tc.want {
			t.Fatalf("Solution(%q) = %q; want %q", tc.in, got, tc.want)
		}
	}
}