package worker

import (
	"context"
	"time"
)

// FirstResult runs fn in the background and returns its result, or an
// error when the context expires first.
func FirstResult(ctx context.Context, fn func() int) (int, error) {
	out := make(chan int)
	go func() {
		out <- fn()
	}()
	select {
	case v := <-out:
		return v, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// WithDeadline is a convenience wrapper around FirstResult.
func WithDeadline(d time.Duration, fn func() int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	return FirstResult(ctx, fn)
}
