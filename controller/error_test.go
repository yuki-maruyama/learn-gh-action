package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_errorHandler(t *testing.T) {
	w := httptest.NewRecorder()
	type args struct {
		w          http.ResponseWriter
		statusCode int
		message    string
	}
	type want struct {
		statusCode int
		message    string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "error",
			args: args{
				w:          w,
				statusCode: 500,
				message:    "error",
			},
			want: want{
				statusCode: 500,
				message:    "error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorHandler(tt.args.w, tt.args.statusCode, tt.args.message)
			if w.Code != tt.want.statusCode {
				t.Errorf("status code want %d, but %d", tt.want.statusCode, w.Code)
			}
			if w.Body.String() != tt.want.message {
				t.Errorf("want %s, but %s", tt.want.message, w.Body.String())
			}
		})
	}
}
