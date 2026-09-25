// Package currency converts human-entered money strings to integer cents.
package currency

import (
	"errors"
	"strconv"
	"strings"
)

// ErrMalformed is returned when the input is not a decimal amount.
var ErrMalformed = errors.New("currency: malformed amount")

// ToCents parses "12.34" into 1234. It rejects negative values and more
// than two decimal places so callers can treat the result as exact.
func ToCents(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" || strings.HasPrefix(s, "-") {
		return 0, ErrMalformed
	}
	whole, frac, _ := strings.Cut(s, ".")
	if len(frac) > 2 {
		return 0, ErrMalformed
	}
	for len(frac) < 2 {
		frac += "0"
	}
	w, err := strconv.ParseInt(whole, 10, 64)
	if err != nil {
		return 0, ErrMalformed
	}
	f, err := strconv.ParseInt(frac, 10, 64)
	if err != nil {
		return 0, ErrMalformed
	}
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
