package database

import (
	"time"
)

type FeedType string

type Feed struct {
	ID           int64
	Name         string
	Url          string
	CreatedAt    *time.Time
	LastSyncedAt *time.Time
}

type FeedWithArticleCount struct {
	Feed
	ArticleCount int
}

func (f Feed) Insert() (int64, error) {
	result, err := db.Exec("INSERT INTO feed (name, url) VALUES (?,?)", f.Name, f.Url)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func FindAllFeedsWithArticleCount() (*[]FeedWithArticleCount, error) {
	rows, err := db.Query("SELECT f.id, f.name, f.url, f.createdAt, f.lastSyncedAt, COUNT(a.id) FROM feed as f LEFT JOIN article as a ON a.feedId = f.id;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feeds []FeedWithArticleCount

	for rows.Next() {
		var feed FeedWithArticleCount
		err := rows.Scan(&feed.ID, &feed.Name, &feed.Url, &feed.CreatedAt, &feed.LastSyncedAt, &feed.ArticleCount)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &feeds, nil
}
