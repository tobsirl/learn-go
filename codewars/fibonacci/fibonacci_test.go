package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Test 1", 0, 0},
		{"Test 2", 1, 1},
		{"Test 3", 2, 1},
		{"Test 4", 3, 2},
		{"Test 5", 4, 3},
		{"Test 6", 5, 5},
		{"Test 7", 6, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}