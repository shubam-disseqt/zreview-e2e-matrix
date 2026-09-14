package search

import (
	"database/sql"
	"fmt"
)

// SearchUsers returns rows matching a case-insensitive name substring.
// BUG 3: SQL injection via fmt.Sprintf string concatenation.
func SearchUsers(db *sql.DB, name string) (*sql.Rows, error) {
	q := fmt.Sprintf("SELECT id, name FROM users WHERE lower(name) LIKE '%%%s%%'", name)
	return db.Query(q)
}
