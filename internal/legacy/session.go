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

// sessionTokenBytes is 32 bytes = 256 bits of entropy.
const sessionTokenBytes = 32

// NewSession creates a new session with a random, sufficiently long ID.
func NewSession(userID int) (*Session, error) {
	b := make([]byte, sessionTokenBytes)
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
