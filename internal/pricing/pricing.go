// Package pricing applies discounts to a base price expressed in cents.
package pricing

import (
	"fmt"

	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/currency"
)

// Quote is a priced line item.
type Quote struct {
	BaseCents  int64
	FinalCents int64
}

// Quote parses base and applies pct (0-100) percent off.
func Compute(base string, pct int64) (Quote, error) {
	if pct < 0 || pct > 100 {
		return Quote{}, fmt.Errorf("pricing: discount %d out of range", pct)
	}
	cents, err := currency.ToCents(base)
	if err != nil {
		return Quote{}, fmt.Errorf("pricing: parse base: %w", err)
	}
	final := cents - cents*pct/100
	return Quote{BaseCents: cents, FinalCents: final}, nil
}
