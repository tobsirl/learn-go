package simplemultiplication

import "testing"

func TestSimpleMultiplication(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   int
	}{
		{"Test 1", 2, 16},
		{"Test 2", 1, 9},
		{"Test 3", 8, 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleMultiplication(tt.number); got != tt.want {
				t.Errorf("SimpleMultiplication(%d) = %d, want %d", tt.number, got, tt.want)
			}
		})
	}
}
