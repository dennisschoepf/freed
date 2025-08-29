package database

import (
	"time"
)

type FeedEntity struct {
	ID           int64
	Name         string
	Url          string
	CreatedAt    *time.Time
	LastSyncedAt *time.Time
}

type FeedEntityWithArticleCount struct {
	FeedEntity
	ArticleCount int
}

func (f FeedEntity) Insert() (int64, error) {
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

// This might be a problem if we have a LOT of feeds configured
// as of now this works, so it is not at the top of the priority
// list
func FindAllFeeds() (*[]FeedEntity, error) {
	rows, err := db.Query("SELECT id, name, url, createdAt, lastSyncedAt FROM feed")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feeds []FeedEntity

	for rows.Next() {
		var feed FeedEntity
		err := rows.Scan(&feed.ID, &feed.Name, &feed.Url, &feed.CreatedAt, &feed.LastSyncedAt)
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

func FindAllFeedsWithArticleCount() (*[]FeedEntityWithArticleCount, error) {
	rows, err := db.Query("SELECT f.id, f.name, f.url, f.createdAt, f.lastSyncedAt, COUNT(a.id) FROM feed as f LEFT JOIN article as a ON a.feedId = f.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feeds []FeedEntityWithArticleCount

	for rows.Next() {
		var feed FeedEntityWithArticleCount
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
