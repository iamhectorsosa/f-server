package testutils

import (
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

func NewMemoryDB() (db *sql.DB, cleanup func() error, err error) {
	migrations := map[string]string{
		"001_schema.sql": `-- +goose Up
CREATE TABLE IF NOT EXISTS players (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS teams (
    id INTEGER PRIMARY KEY,
    player_1 TEXT NOT NULL,
    player_2 TEXT NOT NULL,
    score INTEGER NOT NULL,
    FOREIGN KEY (player_1) REFERENCES players (id) ON DELETE CASCADE,
    FOREIGN KEY (player_2) REFERENCES players (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS matches (
    id INTEGER PRIMARY KEY,
    team_a INTEGER NOT NULL,
    team_b INTEGER NOT NULL,
    FOREIGN KEY (team_a) REFERENCES teams (id) ON DELETE CASCADE,
    FOREIGN KEY (team_b) REFERENCES teams (id) ON DELETE CASCADE
);`,
		"002_seed.sql": `-- +goose Up
INSERT INTO players (id, email) VALUES
('player1', 'player1@example.com'),
('player2', 'player2@example.com'),
('player3', 'player3@example.com'),
('player4', 'player4@example.com');

INSERT INTO teams (player_1, player_2, score) VALUES
('player1', 'player2', 7),
('player3', 'player4', 3);

INSERT INTO matches (team_a, team_b) VALUES
(1, 2);`,
	}

	fs, cleanupFs, err := CreateMigrationsFs("sql/migrations", migrations)
	if err != nil {
		return nil, nil, err
	}
	defer cleanupFs()

	db, err = sql.Open("libsql", ":memory:")
	if err != nil {
		return nil, nil, err
	}
	cleanup = db.Close

	goose.SetBaseFS(fs)

	if err := goose.SetDialect("turso"); err != nil {
		return nil, cleanup, err
	}

	if err := goose.Up(db, "sql/migrations"); err != nil {
		return nil, cleanup, err
	}

	return db, cleanup, nil
}
