package main

import (
	"embed"
	"log"
	"os"

	"github.com/iamhectorsosa/f-server/config"
	"github.com/iamhectorsosa/f-server/internal/database"
	"github.com/iamhectorsosa/f-server/internal/server"
	"github.com/iamhectorsosa/f-server/internal/store"
	"github.com/joho/godotenv"
)

//go:embed sql/migrations/*.sql
var embedMigrations embed.FS

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("err=%v", err)
	}

	db, cleanup, err := database.NewInMemory(embedMigrations)
	if err != nil {
		log.Fatalf("err=%v", err)
	}
	defer cleanup()

	store := store.New(db)

	config := &config.Config{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),
	}

	server := server.New(store, config)
	log.Printf("Listening on http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())

}
