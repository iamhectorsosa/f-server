package server

import (
	"net/http"

	"github.com/iamhectorsosa/f-server/internal/store"
)

func (s *Server) handleMatchesPost(w http.ResponseWriter, r *http.Request) {
	res, err := decode[*store.NewMatch](r)
	if err != nil {
		encode(w, http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	if err := s.store.AddMatch(r.Context(), *res); err != nil {
		encode(w, http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
