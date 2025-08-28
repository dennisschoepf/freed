package internal

import (
	"fmt"
	"freed/internal/database"
)

func GetArticleUrlForToday() (string, error) {
	count, err := database.CountArticlesReadToday()

	if err != nil {
		return "", err
	}

	if count > 0 {
		return "", fmt.Errorf("Already reached maximum number of articles for the day. Come back tomorrow!")
	}

	article, err := database.FindOneUnreadArticle()

	if err != nil {
		return "", err
	}

	if err := article.MarkAsRead(); err != nil {
		return "", err
	}

	return article.Url, nil
}
