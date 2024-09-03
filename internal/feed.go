package feed

import (
	"freed/internal/database"
)

func Add(url string) error {
	f := database.Feed{
		Name:     url,
		Url:      url,
		FeedType: database.FeedType("RSS"),
	}

	if _, err := f.Insert(); err != nil {
		return err
	}

	return nil
}
