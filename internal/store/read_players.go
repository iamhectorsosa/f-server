package store

import (
	"context"
)

type Player = struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Score int    `json:"score"`
}

func (s *Store) GetPlayers(ctx context.Context) ([]Player, error) {
	p, err := s.queries.GetPlayers(ctx)
	if err != nil {
		return nil, err
	}

	var players []Player
	for _, player := range p {
		players = append(players, Player{
			ID:    player.ID,
			Email: player.Email,
			Score: int(player.Score),
		})
	}

	return players, nil
}
