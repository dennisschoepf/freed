package model

type Feed struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required"`
}
