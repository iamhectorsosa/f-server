package store

import (
	"context"
	"errors"
	"strings"

	"github.com/iamhectorsosa/f-server/internal/queries"
)

type NewMatch struct {
	TeamA struct {
		Player1 string `json:"player_1"`
		Player2 string `json:"player_2"`
		Score   int    `json:"score"`
	} `json:"team_a"`
	TeamB struct {
		Player1 string `json:"player_1"`
		Player2 string `json:"player_2"`
		Score   int    `json:"score"`
	} `json:"team_b"`
}

func (m *NewMatch) Valid(ctx context.Context) error {
	problems := make(map[string]string)
	playersSeen := make(map[string]bool)

	players := []struct {
		key, name string
	}{
		{"team_a.player_1", m.TeamA.Player1},
		{"team_a.player_2", m.TeamA.Player2},
		{"team_b.player_1", m.TeamB.Player1},
		{"team_b.player_2", m.TeamB.Player2},
	}

	for _, p := range players {
		if len(p.name) < 5 || len(p.name) > 15 {
			problems[p.key] = "Player ID must be between 5 and 15 characters"
		} else if playersSeen[p.name] {
			problems[p.key] = "Player has already been used"
		} else {
			playersSeen[p.name] = true
		}
	}

	if m.TeamA.Score < 0 {
		problems["team_a.score"] = "Score for Team A cannot be negative"
	}
	if m.TeamB.Score < 0 {
		problems["team_b.score"] = "Score for Team B cannot be negative"
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

func (s *Store) CreateMatch(ctx context.Context, req NewMatch) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queriesTx := s.queries.WithTx(tx)

	teamA, err := queriesTx.CreateTeam(ctx, queries.CreateTeamParams{
		Player1: req.TeamA.Player1,
		Player2: req.TeamA.Player2,
		Score:   req.TeamA.Score,
	})
	if err != nil {
		return err
	}
	teamB, err := queriesTx.CreateTeam(ctx, queries.CreateTeamParams{
		Player1: req.TeamB.Player1,
		Player2: req.TeamB.Player2,
		Score:   req.TeamB.Score,
	})
	if err != nil {
		return err
	}

	err = queriesTx.CreateMatch(ctx, queries.CreateMatchParams{
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
