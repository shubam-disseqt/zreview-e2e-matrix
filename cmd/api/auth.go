package main

import (
	"fmt"
	"net/http"
)

// GreetingHandler returns a personalized greeting to the caller.
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "friend"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", name)
}
