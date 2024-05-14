package database

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*
var dbMigrations embed.FS

func Connect(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filename)

	if err != nil {
		return nil, err
	}

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Applied migrations - Database is ready!\n", n)

	return db, nil
}
