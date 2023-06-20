package config

import "testing"

func TestNewConfigIni(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test config ini",
			args: args{name: "./example.ini"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewConfigIni(tt.args.name)
		})
	}
}
