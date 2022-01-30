package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

func WithRecovery(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			serverErr := recover()

			if serverErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				r := &Response{
					Success: false,
					Err:     fmt.Sprintf("Crash occurred while serving Request -  URL: %s, Body: %s, with error: %s, trace: %s", r.URL, r.Body, serverErr, debug.Stack()),
				}
				resp, _ := json.Marshal(r)
				w.Write(resp)
				return
			}
		}()

		next.ServeHTTP(w, r)
	}
}
