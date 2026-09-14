package legacy

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Retrier retries fn up to MaxAttempts, respecting context cancellation.
type Retrier struct {
	Delay       time.Duration
	MaxAttempts int
}

// ErrGaveUp is returned once a retrier decides to stop.
var ErrGaveUp = errors.New("retrier gave up")

// Do calls fn until it returns nil, ctx is cancelled, or MaxAttempts is hit.
func (r *Retrier) Do(ctx context.Context, fn func() error) error {
	if r.MaxAttempts <= 0 {
		return fmt.Errorf("retrier: MaxAttempts must be > 0")
	}
	var lastErr error
	for i := 0; i < r.MaxAttempts; i++ {
		if err := fn(); err == nil {
			return nil
		} else {
			lastErr = err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(r.Delay):
		}
	}
	return fmt.Errorf("%w: last error: %v", ErrGaveUp, lastErr)
}
