package guuid

import "testing"

func TestGenUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get guuid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenUUID(); got == tt.want {
				t.Errorf("GenUUID() = %v, want %v", got, tt.want)
			} else {
				t.Log(got, len(got))
			}
		})
	}
}
