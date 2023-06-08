package base64

import "testing"

func TestDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Decode",
			args:    args{s: "YWJjZA====="},
			wantErr: true,
			want:    "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Util.Decode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			if got != tt.want {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "abcd",
			},
			want: Base64Util.Encode("abcd"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Util.Encode(tt.args.s); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			} else {
				t.Log(got)
			}
		})
	}
}
