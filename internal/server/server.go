package server

import (
	"net/http"

	"github.com/iamhectorsosa/f-server/config"
	"github.com/iamhectorsosa/f-server/internal/store"
)

type Server struct {
	store  *store.Store
	config *config.Config
	http.Server
}

func New(store *store.Store, config *config.Config) *Server {
	s := new(Server)
	s.store = store
	s.config = config

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", s.handleHealthGet)

	mux.HandleFunc("GET /api/players", s.handlePlayersGet)
	mux.HandleFunc("POST /api/players", s.middlewareProtected(s.handlePlayersPost))

	mux.HandleFunc("POST /api/matches", s.middlewareProtected(s.handleMatchesPost))

	var handler http.Handler = mux
	handler = s.middlewareDevLogger(handler)

	s.Addr = ":" + s.config.Port
	s.Handler = handler

	return s
}
