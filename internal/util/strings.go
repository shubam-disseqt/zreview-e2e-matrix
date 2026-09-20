package util

import "strings"

func TrimAndLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func HasAnyPrefix(s string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
