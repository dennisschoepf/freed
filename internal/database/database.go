package database

import (
	"database/sql"
	"embed"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

var db *sql.DB

//go:embed migrations/*
var dbMigrations embed.FS

func Init(filename string) error {
	dbOptions := "?_fk=on&_journal=WAL&sync=normal"
	db, err := sql.Open("sqlite3", filename+dbOptions)

	if err != nil {
		return err
	}

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		return err
	}

	return db.Ping()
}
