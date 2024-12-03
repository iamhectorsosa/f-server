// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package queries

import (
	"context"
)

const addMatch = `-- name: AddMatch :exec
INSERT INTO matches (team_a, team_b) VALUES (?, ?)
`

type AddMatchParams struct {
	TeamA int `json:"team_a"`
	TeamB int `json:"team_b"`
}

func (q *Queries) AddMatch(ctx context.Context, arg AddMatchParams) error {
	_, err := q.db.ExecContext(ctx, addMatch, arg.TeamA, arg.TeamB)
	return err
}

const addPlayer = `-- name: AddPlayer :exec
INSERT INTO players (id, email) VALUES (?, ?)
`

type AddPlayerParams struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) AddPlayer(ctx context.Context, arg AddPlayerParams) error {
	_, err := q.db.ExecContext(ctx, addPlayer, arg.ID, arg.Email)
	return err
}

const addTeam = `-- name: AddTeam :one
INSERT INTO teams (player_1, player_2, score) VALUES (?, ?, ?) RETURNING id, player_1, player_2, score
`

type AddTeamParams struct {
	Player1 string `json:"player_1"`
	Player2 string `json:"player_2"`
	Score   int    `json:"score"`
}

func (q *Queries) AddTeam(ctx context.Context, arg AddTeamParams) (Team, error) {
	row := q.db.QueryRowContext(ctx, addTeam, arg.Player1, arg.Player2, arg.Score)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Player1,
		&i.Player2,
		&i.Score,
	)
	return i, err
}

const getPlayers = `-- name: GetPlayers :many
SELECT
    p.id,
    p.email,
    CAST(SUM(COALESCE(t.score, 0)) AS INTEGER) AS score
FROM players AS p
LEFT JOIN teams AS t
    ON p.id IN (t.player_1, t.player_2)
GROUP BY p.id
ORDER BY score DESC
`

type GetPlayersRow struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Score int64  `json:"score"`
}

func (q *Queries) GetPlayers(ctx context.Context) ([]GetPlayersRow, error) {
	rows, err := q.db.QueryContext(ctx, getPlayers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPlayersRow
	for rows.Next() {
		var i GetPlayersRow
		if err := rows.Scan(&i.ID, &i.Email, &i.Score); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}