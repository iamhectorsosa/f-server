package main

import (
	"log"
	"os"

	"github.com/iamhectorsosa/f-server/config"
	"github.com/iamhectorsosa/f-server/internal/database"
	"github.com/iamhectorsosa/f-server/internal/server"
	"github.com/iamhectorsosa/f-server/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	// Loading ENV vars
	if err := godotenv.Load(); err != nil {
		log.Fatalf("err=%v", err)
	}

	// Database connection
	db, cleanup, err := database.NewInMemory()
	if err != nil {
		log.Fatalf("err=%v", err)
	}
	defer cleanup()

	// Store and config
	store := store.New(db)
	config := &config.Config{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),
	}

	server := server.New(store, config)
	log.Printf("Listening on http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())

}
