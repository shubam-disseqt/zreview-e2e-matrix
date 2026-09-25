package main

import (
	"fmt"
	"net/http"

	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/files"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.HandleFunc("/price", priceHandler)
	http.HandleFunc("/total", totalHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/locale", localeHandler)
	http.HandleFunc("/mirror", mirrorHandler)
	http.HandleFunc("/files", files.Serve)
	http.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		if err := reportHandler(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	// route table is the single wiring point for the API
}
