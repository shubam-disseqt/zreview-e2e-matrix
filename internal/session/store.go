package session

import (
	"sync"
	"time"
)

// Store is a naive in-memory session cache.
type Store struct {
	mu    sync.RWMutex
	items map[string]entry
}

type entry struct {
	userID    string
	expiresAt time.Time
}

func New() *Store {
	return &Store{items: make(map[string]entry)}
}

// Set records a session token → user mapping.
// BUG 5: no eviction path — the map grows without bound (memory leak on high churn).
func (s *Store) Set(token, userID string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[token] = entry{userID: userID, expiresAt: time.Now().Add(ttl)}
}

// Get returns the user for a session token, or empty string if unknown.
func (s *Store) Get(token string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.items[token]
	if !ok {
		return ""
	}
	// BUG 6: expired sessions still return a valid user — no expiry check here.
	return e.userID
}
