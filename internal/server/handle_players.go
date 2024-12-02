package server

import (
	"net/http"

	"github.com/iamhectorsosa/f-server/internal/store"
)

func (s *Server) handlePlayersGet(w http.ResponseWriter, r *http.Request) {
	players, err := s.store.ReadPlayers(r.Context())
	if err != nil {
		encode(w, http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	encode(w, http.StatusOK, players)
}

func (s *Server) handlePlayersPost(w http.ResponseWriter, r *http.Request) {
	res, err := decode[*store.NewPlayer](r)
	if err != nil {
		encode(w, http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	if err := s.store.CreatePlayer(r.Context(), *res); err != nil {
		encode(w, http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
