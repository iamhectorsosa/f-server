package server

import (
	"log"
	"net/http"

	"github.com/iamhectorsosa/f-server/internal/auth"
)

func (s *Server) middlewareDevLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.config.Env == "development" {
			log.Printf("%s %s", r.Method, r.URL.Path)
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) middlewareProtected(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			encode(w, http.StatusInternalServerError, struct {
				Error string `json:"error"`
			}{
				Error: err.Error(),
			})
			return
		}
		if authToken != s.config.AuthToken {
			encode(w, http.StatusUnauthorized, struct {
				Error string `json:"error"`
			}{
				Error: "Unauthorized",
			})
			return
		}
		next(w, r)
	})
}
