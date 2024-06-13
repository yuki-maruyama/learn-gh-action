package util

import (
	"fmt"
	"testing"
)

var testVal = 128

func TestStrToIntPtr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "success",
			args: args{
				s: fmt.Sprint(testVal),
			},
			want: &testVal,
		},
		{
			name: "failed",
			args: args{
				s: "qwerty",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrToIntPtr(tt.args.s)
			if got == nil && tt.want == nil {
				return
			}
			if got == nil && tt.want != nil {
				t.Errorf("StrToIntPtr() = nil, want %v", *tt.want)
			} else if *got != *tt.want {
				t.Errorf("StrToIntPtr() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
