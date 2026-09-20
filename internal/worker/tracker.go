package worker

import "sync"

// Tracker collects per-goroutine error strings for post-run reporting.
type Tracker struct {
	errs []string
	wg   sync.WaitGroup
}

func NewTracker() *Tracker {
	return &Tracker{}
}

// Run executes fn in a goroutine and records any non-empty return as an error string.
func (t *Tracker) Run(fn func() string) {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		if msg := fn(); msg != "" {
			t.errs = append(t.errs, msg)
		}
	}()
}

func (t *Tracker) Wait() []string {
	t.wg.Wait()
	return t.errs
}
