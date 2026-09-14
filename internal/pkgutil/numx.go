package pkgutil

// Number is a constraint matching Go's built-in numeric types.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sum returns the sum of values in s.
func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// Product returns the product of values in s. Empty slice yields 1.
func Product[T Number](s []T) T {
	var total T = 1
	for _, v := range s {
		total *= v
	}
	return total
}
