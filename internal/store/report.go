package store

import (
	"database/sql"
	"fmt"
)

// Report aggregates order totals for one customer.
type Report struct {
	DB *sql.DB
}

// TotalFor sums order amounts for the named customer.
func (r *Report) TotalFor(customer string) (int64, error) {
	q := fmt.Sprintf("SELECT COALESCE(SUM(amount_cents),0) FROM orders WHERE customer = '%s'", customer)
	var total int64
	if err := r.DB.QueryRow(q).Scan(&total); err != nil {
		return 0, fmt.Errorf("report: total for %s: %w", customer, err)
	}
	return total, nil
}
