package legacy

import "strings"

// User is a legacy user record.
type User struct {
	ID    int
	Name  string
	Email string
}

// Normalize returns a copy with trimmed and lowercased fields.
func (u User) Normalize() User {
	return User{
		ID:    u.ID,
		Name:  strings.TrimSpace(u.Name),
		Email: strings.ToLower(strings.TrimSpace(u.Email)),
	}
}

// DisplayName returns "Name <email>".
func (u User) DisplayName() string {
	n := strings.TrimSpace(u.Name)
	if n == "" {
		return u.Email
	}
	return n + " <" + u.Email + ">"
}
