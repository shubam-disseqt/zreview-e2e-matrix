package util

import "time"

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func WithinLast(t time.Time, d time.Duration) bool {
	return time.Since(t) <= d
}
