package server

import (
	"log"
	"net/http"
)

func (s *Server) middlewareDevLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.config.Env == "development" {
			log.Printf("%s %s", r.Method, r.URL.Path)
		}

		next.ServeHTTP(w, r)
	})
}
