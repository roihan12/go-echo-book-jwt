package repository

import (
	"echo-book/entity"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewBookRepository(t *testing.T) {
	type args struct {
		dbConn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want BookRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookRepository(tt.args.dbConn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_InsertBook(t *testing.T) {
	type args struct {
		b entity.Book
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
		want entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.InsertBook(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookConnection.InsertBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_UpdateBook(t *testing.T) {
	type args struct {
		b entity.Book
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
		want entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.UpdateBook(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookConnection.UpdateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_DeleteBook(t *testing.T) {
	type args struct {
		b entity.Book
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.db.DeleteBook(tt.args.b)
		})
	}
}

func Test_bookConnection_FindBookByID(t *testing.T) {
	type args struct {
		bookID uint64
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
		want entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.FindBookByID(tt.args.bookID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookConnection.FindBookByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_AllBook(t *testing.T) {
	tests := []struct {
		name string
		db   *bookConnection
		want []entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.AllBook(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookConnection.AllBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
