package model

type User struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
