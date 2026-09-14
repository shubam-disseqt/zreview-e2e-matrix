package legacy

import "sync"

// Queue is a simple thread-safe FIFO queue of strings.
type Queue struct {
	mu    sync.Mutex
	items []string
}

// NewQueue returns an empty queue.
func NewQueue() *Queue {
	return &Queue{items: make([]string, 0, 16)}
}

// Push appends an item.
func (q *Queue) Push(item string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Pop removes and returns the head, or "" if empty.
func (q *Queue) Pop() (string, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return "", false
	}
	head := q.items[0]
	q.items = q.items[1:]
	return head, true
}

// Len returns the current length.
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}
