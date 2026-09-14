package legacy

import (
	"database/sql"
	"fmt"
	"net/http"
)

// UserLookup handles user lookup by name from a query param.
type UserLookup struct {
	DB *sql.DB
}

// Handle serves GET /user?name=...
// Uses a parameterized query — no string concatenation with untrusted input.
func (u *UserLookup) Handle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	row := u.DB.QueryRow("SELECT id, email FROM users WHERE name = ?", name)

	var id int
	var email string
	if err := row.Scan(&id, &email); err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%d,%s\n", id, email)
}
