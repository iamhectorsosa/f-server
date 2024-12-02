package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

type Validator interface {
	Valid(ctx context.Context) error
}

func decode[T Validator](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("err=%v", err)
	}
	if err := v.Valid(r.Context()); err != nil {
		return v, fmt.Errorf("err=%v", err)

	}
	return v, nil
}
