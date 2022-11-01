package dto

import "github.com/go-playground/validator/v10"

type RegisterDTO struct {
	Name     string `json:"name" form:"name" validate:"required,min=4,max=15"`
	Email    string `json:"email" form:"email"  validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,omitempty,min=4"`
}

func (input *RegisterDTO) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
