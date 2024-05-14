package model

import "time"

// TODO: Tag for article?
type Article struct {
	ID     int64     `json:"id"`
	Name   string    `json:"name"`
	Url    string    `json:"url" validate:"required,url"`
	Read   bool      `json:"read"`
	ReadAt time.Time `json:"readAt"`
	FeedID *int64    `json:"feedId"`
}
