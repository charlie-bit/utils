package time

import "testing"

func TestGetCurrentDate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentDateYYMM(); got != tt.want {
				t.Errorf("GetCurrentDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
