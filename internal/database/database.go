package database

import (
	"database/sql"
	"embed"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*
var dbMigrations embed.FS

func Connect(path string) (*sql.DB, error) {
	dbOptions := "?_fk=on&_journal=WAL&sync=normal"
	db, err := sql.Open("sqlite3", path+dbOptions)

	if err != nil {
		return nil, err
	}

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
