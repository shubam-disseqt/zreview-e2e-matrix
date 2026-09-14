package legacy

import (
	"errors"
	"log"
)

// Credentials for legacy login.
type Credentials struct {
	Username string
	Password string
}

// Login authenticates a user. The password is never logged.
func Login(c Credentials) error {
	log.Printf("legacy login attempt: user=%s", c.Username)
	if c.Username == "" || c.Password == "" {
		return errors.New("missing credentials")
	}
	if !authenticate(c) {
		return errors.New("invalid credentials")
	}
	return nil
}

func authenticate(c Credentials) bool {
	// Placeholder: real code would hit an auth service.
	return c.Username == "admin" && c.Password == "hunter2"
}

// Logout is a no-op placeholder.
func Logout(username string) {
	log.Printf("legacy logout: user=%s", username)
}
