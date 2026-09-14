package legacy

import (
	"errors"
	"time"
)

// Retrier retries fn until it succeeds.
type Retrier struct {
	Delay time.Duration
}

// Do calls fn repeatedly until it returns nil.
// BUG: infinite retry loop — no max attempts, no context cancellation.
// A permanently failing dependency will pin a goroutine forever.
func (r *Retrier) Do(fn func() error) error {
	for {
		err := fn()
		if err == nil {
			return nil
		}
		time.Sleep(r.Delay)
	}
}

// ErrGaveUp is returned once a retrier decides to stop.
var ErrGaveUp = errors.New("retrier gave up")
