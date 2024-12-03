-- name: AddPlayer :exec
INSERT INTO players (id, email) VALUES (?, ?);

-- name: AddTeam :one
INSERT INTO teams (player_1, player_2, score) VALUES (?, ?, ?) RETURNING *;

-- name: AddMatch :exec
INSERT INTO matches (team_a, team_b) VALUES (?, ?);

-- name: GetPlayers :many
SELECT
    p.id,
    p.email,
    CAST(SUM(COALESCE(t.score, 0)) AS INTEGER) AS score
FROM players AS p
LEFT JOIN teams AS t
    ON p.id IN (t.player_1, t.player_2)
GROUP BY p.id
ORDER BY score DESC;
