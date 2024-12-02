package database

import (
	"database/sql"
	"io/fs"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

func NewInMemory(migrations fs.FS) (db *sql.DB, cleanup func() error, err error) {
	db, err = sql.Open("libsql", ":memory:")
	if err != nil {
		return nil, nil, err
	}
	cleanup = db.Close

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("turso"); err != nil {
		return nil, cleanup, err
	}

	if err := goose.Up(db, "sql/migrations"); err != nil {
		return nil, cleanup, err
	}

	return db, cleanup, nil

}
