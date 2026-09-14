package legacy

import (
	"container/list"
	"sync"
)

// Cache is a bounded LRU string cache.
type Cache struct {
	mu   sync.Mutex
	max  int
	m    map[string]*list.Element
	ll   *list.List
}

type entry struct {
	k, v string
}

// NewCache returns an LRU cache holding at most max entries.
func NewCache(max int) *Cache {
	if max <= 0 {
		max = 1024
	}
	return &Cache{max: max, m: make(map[string]*list.Element, max), ll: list.New()}
}

// Set stores v under k, evicting the oldest entry once full.
func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.m[k]; ok {
		el.Value.(*entry).v = v
		c.ll.MoveToFront(el)
		return
	}
	el := c.ll.PushFront(&entry{k: k, v: v})
	c.m[k] = el
	if c.ll.Len() > c.max {
		old := c.ll.Back()
		if old != nil {
			c.ll.Remove(old)
			delete(c.m, old.Value.(*entry).k)
		}
	}
}

// Get returns the value for k and marks it recently used.
func (c *Cache) Get(k string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	el, ok := c.m[k]
	if !ok {
		return "", false
	}
	c.ll.MoveToFront(el)
	return el.Value.(*entry).v, true
}
