package model

type User struct {
	ID        uint32
	FirstName string `json: "firstName"`
	Email     string `json: "email"`
}
