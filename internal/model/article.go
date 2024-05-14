package model

import "time"

// TODO: Tag for article?
type Article struct {
	ID     int64
	Name   string
	Url    string
	Read   bool
	ReadAt time.Time
	FeedID *int64
}
