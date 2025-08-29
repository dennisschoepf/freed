package database

import (
	"time"
)

type ArticleEntity struct {
	ID     string
	Name   string
	Url    string
	ReadAt *time.Time
	FeedId int64
}

type InsertMultipleOptions struct {
	IgnoreDuplicates bool
}

func (a ArticleEntity) Insert() (int64, error) {
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

func (a ArticleEntity) MarkAsRead() error {
	result, err := db.Exec("UPDATE article SET readAt = datetime() WHERE id = ?;", a.ID)

	if err != nil {
		return err
	}

	if _, err := result.RowsAffected(); err != nil {
		return err
	}

	return nil
}

func InsertMultipleArticles(articles []ArticleEntity) error {
	return insertMultipleArticlesWithOpts(articles, InsertMultipleOptions{IgnoreDuplicates: false})
}

func InsertIgnoreMultipleArticles(articles []ArticleEntity) error {
	return insertMultipleArticlesWithOpts(articles, InsertMultipleOptions{IgnoreDuplicates: true})
}

func insertMultipleArticlesWithOpts(articles []ArticleEntity, opts InsertMultipleOptions) error {
	if len(articles) == 0 {
		return nil
	}

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	insertStmt := "INSERT "

	if opts.IgnoreDuplicates {
		insertStmt = insertStmt + "OR IGNORE "
	}

	finalStmt := insertStmt + "INTO article (name, url, readAt, feedId) VALUES (?, ?, ?, ?)"

	stmt, err := tx.Prepare(finalStmt)

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

func FindOneUnreadArticle() (*ArticleEntity, error) {
	var article ArticleEntity

	row := db.QueryRow("SELECT * FROM article WHERE readAt IS NULL ORDER BY RANDOM() LIMIT 1")
	err := row.Scan(&article.ID, &article.Name, &article.Url, &article.ReadAt, &article.FeedId)

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func CountArticlesReadToday() (int, error) {
	var count int

	result := db.QueryRow("SELECT COUNT(*) FROM article WHERE date(readAt) = date()")
	err := result.Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}
