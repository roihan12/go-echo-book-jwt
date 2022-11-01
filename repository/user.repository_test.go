package repository

import (
	"echo-book/entity"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userConnection_InsertUser(t *testing.T) {
	type args struct {
		user entity.User
	}
	tests := []struct {
		name string
		db   *userConnection
		args args
		want entity.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.InsertUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userConnection.InsertUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userConnection_UpdateUser(t *testing.T) {
	type args struct {
		user entity.User
	}
	tests := []struct {
		name string
		db   *userConnection
		args args
		want entity.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.UpdateUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userConnection.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userConnection_VerifyCredential(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name string
		db   *userConnection
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.VerifyCredential(tt.args.email, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userConnection.VerifyCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userConnection_IsDuplicateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name   string
		db     *userConnection
		args   args
		wantTx *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTx := tt.db.IsDuplicateEmail(tt.args.email); !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("userConnection.IsDuplicateEmail() = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}

func Test_userConnection_FindByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		db   *userConnection
		args args
		want entity.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.FindByEmail(tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userConnection.FindByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userConnection_ProfileUser(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name string
		db   *userConnection
		args args
		want entity.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.ProfileUser(tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userConnection.ProfileUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashAndSalt(t *testing.T) {
	type args struct {
		pwd []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashAndSalt(tt.args.pwd); got != tt.want {
				t.Errorf("hashAndSalt() = %v, want %v", got, tt.want)
			}
		})
	}
}
