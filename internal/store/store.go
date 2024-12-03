package store

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/iamhectorsosa/f-server/internal/database/queries"
)

var validate = validator.New()

type Store struct {
	db      *sql.DB
	queries *queries.Queries
}

func New(db *sql.DB) *Store {
	return &Store{
		db:      db,
		queries: queries.New(db),
	}
}

func (s *Store) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
