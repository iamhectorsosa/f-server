package store

import (
	"context"
)

type Player = struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Score int64  `json:"score"`
}

func (s *Store) ReadPlayers(ctx context.Context) ([]Player, error) {
	p, err := s.queries.ReadPlayers(ctx)
	if err != nil {
		return nil, err
	}
	var players []Player
	for _, player := range p {
		players = append(players, Player{
			ID:    player.ID,
			Email: player.Email,
			Score: player.Score,
		})
	}

	return players, nil
}
