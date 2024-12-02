package store

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"github.com/iamhectorsosa/f-server/internal/queries"
)

type NewPlayer struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (c *NewPlayer) Valid(ctx context.Context) error {
	problems := make(map[string]string)
	if len(c.ID) < 5 || len(c.ID) > 15 {
		problems["id"] = "id cannot be empty, shorter than 5 or longer than 15 characters"
	}

	if _, err := mail.ParseAddress(c.Email); err != nil {
		problems["email"] = "must provide valid email"
	}

	if len(problems) > 0 {
		var err []string
		for field, error := range problems {
			err = append(err, field+": "+error)
		}

		return errors.New(strings.Join(err, "\n"))
	}

	return nil
}

func (s *Store) CreatePlayer(ctx context.Context, req NewPlayer) error {
	return s.queries.CreatePlayer(ctx, queries.CreatePlayerParams{
		ID:    req.ID,
		Email: req.Email,
	})
}
