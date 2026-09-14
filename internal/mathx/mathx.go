package mathx

// Clamp constrains v to [lo, hi]. Panics if lo > hi.
func Clamp(v, lo, hi int) int {
	if lo > hi {
		panic("mathx: lo > hi")
	}
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// AbsDiff returns |a-b|.
func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
