package vaultstore

import "testing"

func TestIsToken(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is token",
			args: args{
				s: "tk_12345",
			},
			want: true,
		},
		{
			name: "is not token",
			args: args{
				s: "tkn_123456",
			},
			want: false,
		},
		{
			name: "is not token",
			args: args{
				s: "12345",
			},
			want: false,
		},
		{
			name: "is not token",
			args: args{
				s: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsToken(tt.args.s); got != tt.want {
				t.Errorf("IsToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
