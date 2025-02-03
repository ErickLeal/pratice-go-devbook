package models

import _ "github.com/go-playground/validator/v10"

type UserUpdateRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=150"`
	Nick  string `json:"nick" validate:"required,min=1,max=150"`
	Email string `json:"email" validate:"required,email"`
}
