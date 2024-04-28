package model

type User struct {
	ID        string `json: id`
	FirstName string `json: "firstName"`
	Email     string `json: "email"`
}
