package worker

import (
	"sync"
	"time"
)

// Pool tracks per-task completion timestamps for observability.
type Pool struct {
	completedAt map[string]time.Time
	wg          sync.WaitGroup
}

func NewPool() *Pool {
	return &Pool{completedAt: make(map[string]time.Time)}
}

// Submit runs fn asynchronously and records the completion time keyed by id.
func (p *Pool) Submit(id string, fn func()) {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		fn()
		p.completedAt[id] = time.Now()
	}()
}

func (p *Pool) Wait() { p.wg.Wait() }

func (p *Pool) CompletedAt(id string) (time.Time, bool) {
	t, ok := p.completedAt[id]
	return t, ok
}
