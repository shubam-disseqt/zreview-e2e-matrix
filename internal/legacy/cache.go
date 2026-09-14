package legacy

import "sync"

// Cache is a naive in-memory string cache.
// BUG: unbounded map — no eviction, no max size => memory leak.
// Every unique key held forever for the life of the process.
type Cache struct {
	mu sync.RWMutex
	m  map[string]string
}

// NewCache returns an empty cache.
func NewCache() *Cache {
	return &Cache{m: make(map[string]string)}
}

// Set stores v under k.
func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

// Get returns the value for k.
func (c *Cache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[k]
	return v, ok
}
