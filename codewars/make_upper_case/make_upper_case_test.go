package makeuppercase

import "testing"

func TestMakeUpperCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"hello"}, "HELLO"},
		{"test2", args{"hello world"}, "HELLO WORLD"},
		{"test3", args{"hello world!"}, "HELLO WORLD!"},
	}
	for _, tt := range tests {
		if got := MakeUpperCase(tt.args.str); got != tt.want {
			t.Errorf("%q. MakeUpperCase() = %v, want %v", tt.name, got, tt.want)
		}
	}
}