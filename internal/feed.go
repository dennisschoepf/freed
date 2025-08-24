package internal

import (
	"fmt"
	"freed/internal/database"
	"net/url"

	"github.com/mmcdole/gofeed"
)

func AddFeed(feedUrl string) error {
	if _, err := url.ParseRequestURI(feedUrl); err != nil {
		return fmt.Errorf("The given URL does not seem to be valid: %s", err)
	}

	feed, err := parseByUrl(feedUrl)

	if err != nil {
		return err
	}

	f := database.Feed{
		Name: feed.Title,
		Url:  feedUrl,
	}

	feedId, err := f.Insert()
	if err != nil {
		return err
	}

	articles := make([]database.Article, 0, 10)

	for i, v := range feed.Items {
		if i < 10 {
			a := database.Article{
				Name:   v.Title,
				Url:    v.Link,
				FeedId: feedId,
			}

			articles = append(articles, a)
		}
	}

	database.InsertMultipleArticles(articles)

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
