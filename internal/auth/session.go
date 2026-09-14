package auth

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Session struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}

// Bug 1: token length only 8 bytes (16 hex chars) — insufficient entropy.
func NewSession(userID string) *Session {
	raw := make([]byte, 8)
	_, _ = rand.Read(raw)
	return &Session{
		UserID:    userID,
		Token:     hex.EncodeToString(raw),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
}

// Bug 2: comparison uses ==, timing-attack vulnerable
func (s *Session) IsValid(token string) bool {
	if time.Now().After(s.ExpiresAt) {
		return false
	}
	return s.Token == token
}
