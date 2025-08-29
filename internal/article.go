package internal

import (
	"fmt"
	"time"

	"freed/internal/database"
)

type Article struct {
	ID     string
	Name   string
	Url    string
	ReadAt *time.Time
	FeedId int64
}

func GetArticleUrlForToday() (*Article, error) {
	count, err := database.CountArticlesReadToday()

	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, fmt.Errorf("Already reached maximum number of articles for the day. Come back tomorrow!")
	}

	articleEntity, err := database.FindOneUnreadArticle()

	if err != nil {
		return nil, err
	}

	// For now doing it at this point is enough
	// In the future it might be better to gather user input
	// on if the article was read
	if err := articleEntity.MarkAsRead(); err != nil {
		return nil, err
	}

	article := &Article{
		ID:     articleEntity.ID,
		Name:   articleEntity.Name,
		Url:    articleEntity.Url,
		ReadAt: articleEntity.ReadAt,
		FeedId: articleEntity.FeedId,
	}

	return article, nil
}
