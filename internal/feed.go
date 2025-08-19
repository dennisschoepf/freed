package feed

import (
	"fmt"
	"freed/internal/database"
	"net/url"

	"github.com/mmcdole/gofeed"
)

func Add(feedUrl string) error {
	if _, err := url.ParseRequestURI(feedUrl); err != nil {
		return fmt.Errorf("The given URL does not seem to be valid: %s", err)
	}

	feed, err := parseByUrl(feedUrl)

	if err != nil {
		return err
	}

	fmt.Printf("feed: %+v", feed)

	// TODO: I have everything here (feed), might as well store it right away
	// So, implement a loop to do that here

	f := database.Feed{
		Name: feed.Title,
		Url:  feedUrl,
	}

	if _, err := f.Insert(); err != nil {
		return err
	}

	return nil
}

func parseByUrl(u string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(u)

	if err != nil {
		return nil, err
	}

	return feed, nil
}
