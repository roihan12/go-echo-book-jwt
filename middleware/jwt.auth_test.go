package middleware

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAuthorizeJWT(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthorizeJWT(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
