package pkgutil

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

// RandInt returns a cryptographically random int64 in [0, n). Panics if n <= 0.
func RandInt(n int64) (int64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("pkgutil: RandInt n must be > 0, got %d", n)
	}
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, fmt.Errorf("pkgutil: RandInt read: %w", err)
	}
	// Mask off the sign bit, then take modulo n. Bias is negligible for n << 2^63.
	v := int64(binary.BigEndian.Uint64(b[:]) & 0x7FFFFFFFFFFFFFFF)
	return v % n, nil
}
