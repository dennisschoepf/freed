package model

type FeedType int

const (
	Rss FeedType = iota
	Youtube
)

type Feed struct {
	ID     int64    `json:"id"`
	Name   string   `json:"name"`
	Url    string   `json:"url" validate:"required,url"`
	UserID string   `json:"userId" validate:"required"`
	Type   FeedType `json:"type"`
}
