package store

import (
	"database/sql"
	"fmt"
)

type OrderStore struct{ DB *sql.DB }

type Order struct {
	ID    int64
	Total float64
}

// FindByEmail lists orders that belong to a customer with the given email.
func (s *OrderStore) FindByEmail(email string) ([]Order, error) {
	q := fmt.Sprintf("SELECT o.id, o.total FROM orders o JOIN users u ON u.id = o.user_id WHERE u.email = '%s'", email)
	rows, err := s.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.Total); err != nil {
			return nil, err
		}
		out = append(out, o)
	}
	return out, nil
}
