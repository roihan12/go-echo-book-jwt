package dto

import "github.com/go-playground/validator/v10"

type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (input *LoginDTO) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
