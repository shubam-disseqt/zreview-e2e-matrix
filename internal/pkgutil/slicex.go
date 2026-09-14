package pkgutil

// SliceContains reports whether v is present in s.
func SliceContains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

// SliceMap returns a new slice with fn applied to each element.
func SliceMap[T, U any](s []T, fn func(T) U) []U {
	out := make([]U, len(s))
	for i, v := range s {
		out[i] = fn(v)
	}
	return out
}
