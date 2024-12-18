package store_test

import (
	"context"
	"testing"

	"github.com/iamhectorsosa/f-server/internal/database"
	"github.com/iamhectorsosa/f-server/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestAddMatch(t *testing.T) {
	db, cleanup, err := database.NewInMemory()
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	defer cleanup()
	s := store.New(db)

	newMatch := store.NewMatch{
		TeamA: store.NewTeam{
			Player1: "player1",
			Player2: "player2",
			Score:   1,
		},
		TeamB: store.NewTeam{
			Player1: "player3",
			Player2: "player4",
			Score:   1,
		},
	}

	err = s.AddMatch(context.Background(), newMatch)
	if err != nil {
		t.Fatalf("err=%v", err)
	}

	players, err := s.GetPlayers(context.Background())

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
