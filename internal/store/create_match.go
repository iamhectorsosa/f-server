package store

import (
	"context"
	"errors"
	"strings"

	"github.com/iamhectorsosa/f-server/internal/database/queries"
)

type NewTeam struct {
	Player1 string `json:"player_1" validate:"required,len=5|len=15"`
	Player2 string `json:"player_2" validate:"required,len=5|len=15"`
	Score   int    `json:"score" validate:"gte=0"`
}

type NewMatch struct {
	TeamA NewTeam `json:"team_a"`
	TeamB NewTeam `json:"team_b"`
}

func (m *NewMatch) Valid(ctx context.Context) error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	players := []string{
		m.TeamA.Player1, m.TeamA.Player2,
		m.TeamB.Player1, m.TeamB.Player2,
	}

	playersSeen := make(map[string]bool)
	var problems []string

	for _, player := range players {
		if len(player) == 0 {
			continue
		}
		if playersSeen[player] {
			problems = append(problems, "Player "+player+" has already been used")
		} else {
			playersSeen[player] = true
		}
	}

	if len(problems) > 0 {
		return errors.New(strings.Join(problems, "\n"))
	}

	return nil
}

func (s *Store) AddMatch(ctx context.Context, req NewMatch) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queriesTx := s.queries.WithTx(tx)

	teamA, err := queriesTx.AddTeam(ctx, queries.AddTeamParams{
		Player1: req.TeamA.Player1,
		Player2: req.TeamA.Player2,
		Score:   req.TeamA.Score,
	})
	if err != nil {
		return err
	}
	teamB, err := queriesTx.AddTeam(ctx, queries.AddTeamParams{
		Player1: req.TeamB.Player1,
		Player2: req.TeamB.Player2,
		Score:   req.TeamB.Score,
	})
	if err != nil {
		return err
	}

	err = queriesTx.AddMatch(ctx, queries.AddMatchParams{
		TeamA: teamA.ID,
		TeamB: teamB.ID,
	})
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
