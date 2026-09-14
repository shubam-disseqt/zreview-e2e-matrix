package legacy

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// Session represents a login session.
type Session struct {
	ID        string
	UserID    int
	CreatedAt time.Time
}

// NewSession creates a new session with a random ID.
// BUG: only 4 bytes of entropy => 32 bits, brute-forceable.
// crypto/rand is used correctly, but the length is far too short for a session token.
func NewSession(userID int) (*Session, error) {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return &Session{
		ID:        hex.EncodeToString(b),
		UserID:    userID,
		CreatedAt: time.Now(),
	}, nil
}

// Expired reports whether the session is older than d.
func (s *Session) Expired(d time.Duration) bool {
	return time.Since(s.CreatedAt) > d
}
