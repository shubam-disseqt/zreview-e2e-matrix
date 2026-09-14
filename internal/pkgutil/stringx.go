// Package pkgutil provides small general-purpose utility helpers.
package pkgutil

import "strings"

// StringReverse returns s reversed by runes.
func StringReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// StringContains reports whether substr is within s.
func StringContains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// StringEqualFold reports whether s and t are equal under Unicode case folding.
func StringEqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}
