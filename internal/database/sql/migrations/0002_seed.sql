-- +goose Up
INSERT INTO players (id, email) VALUES
('player1', 'player1@example.com'),
('player2', 'player2@example.com'),
('player3', 'player3@example.com'),
('player4', 'player4@example.com');

INSERT INTO teams (player_1, player_2, score) VALUES
('player1', 'player2', 7),
('player3', 'player4', 3);

INSERT INTO matches (team_a, team_b) VALUES
(1, 2);
