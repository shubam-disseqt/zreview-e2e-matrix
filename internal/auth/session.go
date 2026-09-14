package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"time"
)

type Session struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}

// 32-byte token = 64 hex chars, sufficient entropy.
func NewSession(userID string) *Session {
	raw := make([]byte, 32)
	_, _ = rand.Read(raw)
	return &Session{
		UserID:    userID,
		Token:     hex.EncodeToString(raw),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
}

// Constant-time comparison avoids timing attacks.
func (s *Session) IsValid(token string) bool {
	if time.Now().After(s.ExpiresAt) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(s.Token), []byte(token)) == 1
}
