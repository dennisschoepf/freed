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
	dbOptions := "?_fk=on&_journal=WAL&sync=normal"
	db, err := sql.Open("sqlite3", filename+dbOptions)

	if err != nil {
		return nil, err
	}

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	_, migrateErr := migrate.Exec(db, "sqlite3", migrations, migrate.Up)

	if migrateErr != nil {
		return nil, err
	}

	fmt.Println("Applied migrations - Database is ready!")

	return db, nil
}
