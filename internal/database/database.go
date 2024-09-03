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

func Open(path string) error {
	var err error
	dbOptions := "?_fk=on&_journal=WAL&sync=normal"
	db, err = sql.Open("sqlite3", path+dbOptions)

	if err != nil {
		return err
	}

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	if _, migrateErr := migrate.Exec(db, "sqlite3", migrations, migrate.Up); migrateErr != nil {
		return migrateErr
	}

	if pingErr := db.Ping(); pingErr != nil {
		return pingErr
	}

	return nil
}
