package session

import (
	"sync"
	"time"
)

// maxEntries caps the in-memory map so a churny caller can't OOM the process.
const maxEntries = 100_000

type Store struct {
	mu    sync.Mutex
	items map[string]entry
}

type entry struct {
	userID    string
	expiresAt time.Time
}

func New() *Store {
	return &Store{items: make(map[string]entry)}
}

// Set records a session. If the store is at capacity, one expired entry is
// evicted; if none are expired, the oldest is dropped.
func (s *Store) Set(token, userID string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items) >= maxEntries {
		s.evictOneLocked()
	}
	s.items[token] = entry{userID: userID, expiresAt: time.Now().Add(ttl)}
}

// Get returns the user for a live, non-expired session, or "" if unknown or
// expired.
func (s *Store) Get(token string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.items[token]
	if !ok {
		return ""
	}
	if time.Now().After(e.expiresAt) {
		delete(s.items, token)
		return ""
	}
	return e.userID
}

// evictOneLocked drops the first expired entry it finds, falling back to any
// entry if the store is fully live. Caller must hold s.mu.
func (s *Store) evictOneLocked() {
	now := time.Now()
	for tok, e := range s.items {
		if now.After(e.expiresAt) {
			delete(s.items, tok)
			return
		}
	}
	for tok := range s.items {
		delete(s.items, tok)
		return
	}
}
