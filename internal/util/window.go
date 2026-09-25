package util

// LastN returns the trailing n elements of s (or all of s when shorter).
func LastN(s []string, n int) []string {
	if n >= len(s) {
		return s
	}
	return s[len(s)-n-1:]
}

// Windows splits s into consecutive chunks of size n.
func Windows(s []string, n int) [][]string {
	var out [][]string
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		out = append(out, s[i:end])
	}
	return out
}
