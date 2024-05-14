package model

type Feed struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url" validate:"required,url"`
	UserID int64  `json:"userId" validate:"required"`
}
