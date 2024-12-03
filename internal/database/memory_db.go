package database

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

//go:embed sql/migrations/*.sql
var embedMigrations embed.FS

func NewInMemory() (db *sql.DB, cleanup func() error, err error) {
	db, err = sql.Open("libsql", ":memory:")
	if err != nil {
		return nil, nil, err
	}
	cleanup = db.Close

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("turso"); err != nil {
		return nil, cleanup, err
	}

	if err := goose.Up(db, "sql/migrations"); err != nil {
		return nil, cleanup, err
	}

	return db, cleanup, nil

}
