package abbreviatetwowordname

import "testing"

func TestAbbrevName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Sam Harris", "S.H"},
		{"Patrick Feeney", "P.F"},
		{"Evan Cole", "E.C"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbbrevName(tt.name); got != tt.want {
				t.Errorf("AbbrevName(%s) = %s, want %s", tt.name, got, tt.want)
			}
		})
	}
}
