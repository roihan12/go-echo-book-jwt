package dto

import "github.com/go-playground/validator/v10"

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

// type UserCreatDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
// }

func (input *UserUpdateDTO) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
