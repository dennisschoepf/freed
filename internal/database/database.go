package database

import (
	"embed"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

var DB *sqlx.DB

//go:embed migrations/*
var dbMigrations embed.FS

func Connect(filename string) error {
	DB, err := sqlx.Open("sqlite3", filename)

	if err != nil {
		return err
	}

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	n, err := migrate.Exec(DB.DB, "sqlite3", migrations, migrate.Up)

	if err != nil {
		return err
	}

	fmt.Printf("Applied %d migrations - Database is ready!\n", n)

	return nil
}
