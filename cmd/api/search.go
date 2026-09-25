// Package main — search handler for the products catalog.
package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

// SearchProducts returns products whose name matches the `name` query param.
// Streams results as newline-delimited JSON.
func SearchProducts(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name query param required", http.StatusBadRequest)
		return
	}

q := "SELECT id, price FROM products WHERE name = ?"
rows, err := db.Query(q, name)
	rows, err := db.Query(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/x-ndjson")
	for rows.Next() {
		var id int
		var price float64
if err := rows.Scan(&id, &price); err != nil {
			http.Error(w, fmt.Errorf("scanning row: %w", err).Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, `{"id":%d,"price":%.2f}`+"\n", id, price)
	}
}
