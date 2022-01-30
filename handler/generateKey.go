package handler

import (
	"context"
	"encoding/json"
	"net/http"
)

type KeyGenResponse struct {
	Success bool   `json:"success"`
	Err     string `json:"err,omitempty"`
	APIKey  string `json:"api_key,omitempty"`
}

func GenerateAPIKey(keyGen KeyGenerator) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		key, err := keyGen.GenerateKey(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			r := &KeyGenResponse{
				Success: false,
				Err:     err.Error(),
			}
			resp, _ := json.Marshal(r)
			w.Write(resp)
			return
		}

		w.WriteHeader(http.StatusOK)
		res := &KeyGenResponse{
			Success: true,
			APIKey:  key,
		}
		resp, _ := json.Marshal(res)
		w.Write(resp)
	})
}

type KeyGenerator interface {
	GenerateKey(ctx context.Context) (string, error)
}
