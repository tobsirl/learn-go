package lostwithoutamap

import "testing"

func TestMaps(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3}}, []int{2, 4, 6}},
		{"test2", args{[]int{4, 1, 1, 1, 4}}, []int{8, 2, 2, 2, 8}},
		{"test3", args{[]int{1, 2, 3, 4, 5}}, []int{2, 4, 6, 8, 10}},
	}
	for _, tt := range tests {
		if got := Maps(tt.args.x); !equal(got, tt.want) {
			t.Errorf("%q. Maps() = %v, want %v", tt.name, got, tt.want)
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