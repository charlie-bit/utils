package time

import "testing"

func TestGetCurrentDate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "output 2023/06/08 16:23:20 format time",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentDateYYSlMMSlDD(); got == tt.want {
				t.Errorf("GetCurrentDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
