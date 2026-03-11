package basicvariableassignment

import "testing"

func TestNamevar(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test 1", "codewa.rs"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Namevar(); got != tt.want {
				t.Errorf("Namevar() = %s, want %s", got, tt.want)
			}
		})
	}
}