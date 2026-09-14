package search

import (
	"database/sql"
)

// SearchUsers returns rows matching a case-insensitive name substring.
func SearchUsers(db *sql.DB, name string) (*sql.Rows, error) {
	return db.Query(
		"SELECT id, name FROM users WHERE lower(name) LIKE '%' || lower($1) || '%'",
		name,
	)
}
