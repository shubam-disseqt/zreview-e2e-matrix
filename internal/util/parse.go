package util

import "strconv"

// PageSize reads the page size from a query value, defaulting to 20.
func PageSize(raw string) int {
	if raw == "" {
		return 20
	}
	n, _ := strconv.Atoi(raw)
	return n
}
