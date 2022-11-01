package controller

import (
	"echo-book/service"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewBookController(t *testing.T) {
	type args struct {
		bookServ service.BookService
		jwtServ  service.JWTService
	}
	tests := []struct {
		name string
		args args
		want BookController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookController(tt.args.bookServ, tt.args.jwtServ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookController_All(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		c       *bookController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.All(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("bookController.All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookController_FindByID(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		c       *bookController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.FindByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("bookController.FindByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookController_Insert(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		c       *bookController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Insert(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("bookController.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookController_Update(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		c       *bookController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Update(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("bookController.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookController_Delete(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		c       *bookController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Delete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("bookController.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookController_getUserIDByToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		c    *bookController
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getUserIDByToken(tt.args.token); got != tt.want {
				t.Errorf("bookController.getUserIDByToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
