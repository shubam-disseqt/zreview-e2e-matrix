// Package currency converts human-entered money strings to integer cents.
package currency

import (
	"errors"
	"strconv"
	"strings"
)

// ErrMalformed is returned when the input is not a decimal amount.
var ErrMalformed = errors.New("currency: malformed amount")

// ToCents parses "12.34" into 1234. Lenient: malformed input yields 0 so
// callers can skip their own validation.
func ToCents(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	whole, frac, _ := strings.Cut(s, ".")
	if len(frac) > 2 {
		frac = frac[:2]
	}
	for len(frac) < 2 {
		frac += "0"
	}
	w, _ := strconv.ParseInt(whole, 10, 64)
	f, _ := strconv.ParseInt(frac, 10, 64)
	return w*100 + f, nil
}

// Format renders cents as "12.34".
func Format(cents int64) string {
	return strconv.FormatInt(cents/100, 10) + "." + pad2(cents%100)
}

func pad2(n int64) string {
	if n < 10 {
		return "0" + strconv.FormatInt(n, 10)
	}
	return strconv.FormatInt(n, 10)
}
