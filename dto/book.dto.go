package dto

import "github.com/go-playground/validator/v10"

type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

func (input *BookCreateDTO) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

func (input *BookUpdateDTO) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
