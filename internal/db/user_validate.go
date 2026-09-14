package db

import "strings"

// ValidEmail returns true if e has a single @ and a non-empty local/domain part.
func ValidEmail(e string) bool {
	i := strings.IndexByte(e, '@')
	if i <= 0 || i == len(e)-1 {
		return false
	}
	return strings.IndexByte(e[i+1:], '@') == -1
}
