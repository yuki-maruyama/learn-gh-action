package util

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("FOO", "1")
	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "read env",
			args: args{
				key:      "FOO",
				fallback: "null",
			},
			want: "1",
		},
		{
			name: "fallback",
			args: args{
				key:      "QWERTYUIOP",
				fallback: "fallback",
			},
			want: "fallback",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
