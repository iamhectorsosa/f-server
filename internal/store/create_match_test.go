package store_test

import (
	"context"
	"testing"

	"github.com/iamhectorsosa/f-server/internal/store"
	"github.com/iamhectorsosa/f-server/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestCreateMatch(t *testing.T) {
	db, cleanup, err := testutils.NewMemoryDB()
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	defer cleanup()
	s := store.New(db)

	newMatch := store.NewMatch{
		TeamA: struct {
			Player1 string `json:"player_1"`
			Player2 string `json:"player_2"`
			Score   int    `json:"score"`
		}{
			Player1: "player1",
			Player2: "player2",
			Score:   1,
		},
		TeamB: struct {
			Player1 string `json:"player_1"`
			Player2 string `json:"player_2"`
			Score   int    `json:"score"`
		}{
			Player1: "player3",
			Player2: "player4",
			Score:   1,
		},
	}

	err = s.CreateMatch(context.Background(), newMatch)
	if err != nil {
		t.Fatalf("err=%v", err)
	}

	players, err := s.ReadPlayers(context.Background())

	if err != nil {
		t.Fatalf("err=%v", err)
	}

	assert.NotNil(t, players, "players should not be nil")
	assert.Len(t, players, 4, "expected 4 players")

	expectedPlayers := []store.Player{
		{ID: "player1", Email: "player1@example.com", Score: 8},
		{ID: "player2", Email: "player2@example.com", Score: 8},
		{ID: "player3", Email: "player3@example.com", Score: 4},
		{ID: "player4", Email: "player4@example.com", Score: 4},
	}

	assert.ElementsMatch(t, players, expectedPlayers, "players data should match")
}
