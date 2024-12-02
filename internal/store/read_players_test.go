package store_test

import (
	"context"
	"testing"

	"github.com/iamhectorsosa/f-server/internal/store"
	"github.com/iamhectorsosa/f-server/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestReadPlayers(t *testing.T) {
	db, cleanup, err := testutils.NewMemoryDB()
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	defer cleanup()
	s := store.New(db)

	players, err := s.ReadPlayers(context.Background())

	if err != nil {
		t.Fatalf("err=%v", err)
	}

	assert.NotNil(t, players, "players should not be nil")
	assert.Len(t, players, 4, "expected 4 players")

	expectedPlayers := []store.Player{
		{ID: "player1", Email: "player1@example.com", Score: 7},
		{ID: "player2", Email: "player2@example.com", Score: 7},
		{ID: "player3", Email: "player3@example.com", Score: 3},
		{ID: "player4", Email: "player4@example.com", Score: 3},
	}

	assert.ElementsMatch(t, players, expectedPlayers, "players data should match")
}
