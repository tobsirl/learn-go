package oppositenumber

import "testing"

func TestOpposite(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{1}, -1},
		{"Test 2", args{-5}, 5},
		{"Test 3", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Opposite(tt.args.number); got != tt.want {
				t.Errorf("Opposite() = %v, want %v", got, tt.want)
			}
		})
	}
}