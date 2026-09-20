package main

import (
	"fmt"
	"net/http"
)

// ProfileHandler returns the caller's profile page including their public bio.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	bio := r.URL.Query().Get("bio")
	if bio == "" {
		bio = "(no bio yet)"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<html><body><section><h1>Profile</h1><p>%s</p></section></body></html>", bio)
}
