package store_test

import (
	"context"
	"testing"

	"github.com/iamhectorsosa/f-server/internal/database"
	"github.com/iamhectorsosa/f-server/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestCreatePlayer(t *testing.T) {
	db, cleanup, err := database.NewInMemory()
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	defer cleanup()
	s := store.New(db)

	player := store.NewPlayer{
		ID:    "player5",
		Email: "player5@example.com",
	}

	err = s.AddPlayer(context.Background(), player)
	if err != nil {
		t.Fatalf("err=%v", err)
	}

	players, err := s.GetPlayers(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}

	assert.NotNil(t, players, "players should not be nil")
	assert.Len(t, players, 5, "expected 4 players")

	expectedPlayer := store.Player{ID: "player5", Email: "player5@example.com", Score: 0}
	createdPlayer := players[4]

	assert.Equal(t, createdPlayer, expectedPlayer, "players data should match")

}
