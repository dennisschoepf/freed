package database

import (
	"time"
)

type Article struct {
	Name   string
	Url    string
	ReadAt *time.Time
	FeedId int64
}

func (a Article) Insert() (int64, error) {
	result, err := db.Exec("INSERT INTO article (name, url, readAt, feedId) VALUES (?, ?, ?, ?)", a.Name, a.Url, a.ReadAt, a.FeedId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func InsertMultipleArticles(articles []Article) error {
	if len(articles) == 0 {
		return nil
	}

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO article (name, url, readAt, feedId) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, article := range articles {
		_, err := stmt.Exec(article.Name, article.Url, article.ReadAt, article.FeedId)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
