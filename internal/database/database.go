package database

import (
	"embed"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*
var dbMigrations embed.FS

func New(filename string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", filename)

	if err != nil {
		return nil, err
	}

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	n, err := migrate.Exec(db.DB, "sqlite3", migrations, migrate.Up)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Applied %d migrations - Database is ready!\n", n)

	return db, err
}
