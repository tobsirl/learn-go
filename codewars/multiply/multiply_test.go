package multiply

import "testing"

func TestMultiply(t *testing.T) {
  tests := []struct {
    a, b int
    want int
  }{
    {1, 1, 1},
    {2, 5, 10},
    {5, 10, 50},
    {5, 0, 0},
    {0, 5, 0},
    {0, 0, 0},
  }

  for _, tc := range tests {
    got := Multiply(tc.a, tc.b)
    if got != tc.want {
      t.Fatalf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
    }
  }
}