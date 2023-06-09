package glogrus

import "testing"

func TestSetup(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test log",
			args: args{
				mode: "debug",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Setup(tt.args.mode)
			Warn("debug")
			Error("error")
			Info("info")
		})
	}
}
