package pkgutil

import "time"

// Since returns the duration elapsed since t. Wrapper for symmetry with Until.
func Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Until returns the duration until t.
func Until(t time.Time) time.Duration {
	return time.Until(t)
}
