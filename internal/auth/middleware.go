package auth

import (
	"net/http"
)

var currentSession *Session // package-global — dangerous in real code but not a "bug" per se

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if currentSession == nil || !currentSession.IsValid(token) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
