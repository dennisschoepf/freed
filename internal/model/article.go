package model

import "time"

type Article struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Url    string    `json:"url"`
	Read   bool      `json:"read"`
	ReadAt time.Time `json:"readAt"`
}
