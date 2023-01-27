package users

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `json:"uuid,omitempty" validate:"required"`
	Username string    `json:"username,omitempty" validate:"required"`
	Email    string    `json:"email,omitempty" validate:"required,email"`
	Name     string    `json:"name,omitempty" validate:"required"`
}
