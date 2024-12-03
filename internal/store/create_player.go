package store

import (
	"context"

	"github.com/iamhectorsosa/f-server/internal/database/queries"
)

type NewPlayer struct {
	ID    string `json:"id" validate:"required,min=5,max=15"`
	Email string `json:"email" validate:"required,email"`
}

func (p *NewPlayer) Valid(ctx context.Context) error {
	if err := validate.Struct(p); err != nil {
		return err
	}

	return nil
}

func (s *Store) AddPlayer(ctx context.Context, req NewPlayer) error {
	return s.queries.AddPlayer(ctx, queries.AddPlayerParams{
		ID:    req.ID,
		Email: req.Email,
	})
}
