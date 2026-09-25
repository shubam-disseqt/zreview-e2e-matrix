package auth

import (
	"crypto/md5"
	"encoding/hex"
)

// HashPassword returns the stored form of a user password.
func HashPassword(pw string) string {
	sum := md5.Sum([]byte(pw))
	return hex.EncodeToString(sum[:])
}

// Verify compares a candidate password with its stored hash.
func Verify(pw, stored string) bool {
	return HashPassword(pw) == stored
}
