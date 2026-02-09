package stringrepeat

import "testing"

func TestRepeatStr(t *testing.T) {
	type args struct {
		repetitions int
		value       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{3, "abc"}, "abcabcabc"},
		{"Test 2", args{5, "Hello"}, "HelloHelloHelloHelloHello"},
	}
	for _, tt := range tests {
		if got := RepeatStr(tt.args.repetitions, tt.args.value); got != tt.want {
			t.Errorf("%s: RepeatStr() = %v, want %v", tt.name, got, tt.want)
		}
	}
}