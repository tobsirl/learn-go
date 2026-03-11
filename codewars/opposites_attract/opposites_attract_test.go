package oppositesattract

import "testing"

func TestLovefunc(t *testing.T) {
	type args struct {
		flower1 int
		flower2 int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{1, 4}, true},
		{"test2", args{2, 2}, false},
		{"test3", args{0, 1}, true},
		{"test4", args{0, 0}, false},
	}
	for _, tt := range tests {
		if got := Lovefunc(tt.args.flower1, tt.args.flower2); got != tt.want {
			t.Errorf("%q. Lovefunc() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
