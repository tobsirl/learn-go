package scrabblescore

import "testing"

func TestScrabbleScore(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"STREET", 6},
		{"st re et", 6},
		{"ca bba g  e", 14},
		{"cabbage", 14},
	}
	for _, tt := range tests {
		if got := ScrabbleScore(tt.input); got != tt.expected {
			t.Errorf("ScrabbleScore(%q) = %d, want %d", tt.input, got, tt.expected)
		}
	}
}

