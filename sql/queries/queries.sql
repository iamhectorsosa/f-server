-- name: CreatePlayer :exec
INSERT INTO players (id, email) VALUES (?, ?);

-- name: CreateTeam :one
INSERT INTO teams (player_1, player_2, score) VALUES (?, ?, ?) RETURNING *;

-- name: CreateMatch :exec
INSERT INTO matches (team_a, team_b) VALUES (?, ?);

-- name: ReadPlayers :many
SELECT
    p.id,
    p.email,
    CAST(COALESCE(SUM(t.score), 0) AS INTEGER) AS score
FROM players AS p
LEFT JOIN teams AS t
    ON p.id IN (t.player_1, t.player_2)
GROUP BY p.id
ORDER BY score DESC;
