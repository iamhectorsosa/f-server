package server

import (
	"net/http"
)

func (s *Server) handleHealthGet(w http.ResponseWriter, r *http.Request) {
	encode(w, http.StatusOK, map[string]string{"status": "OK"})
}
