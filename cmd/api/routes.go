package main

import "net/http"

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/greeting", GreetingHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
