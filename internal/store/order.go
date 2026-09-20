package store

import (
	"context"
	"database/sql"
)

type Order struct {
	ID     int64
	UserID int64
	Total  float64
}

type OrderStore struct{ DB *sql.DB }

func (s *OrderStore) Get(ctx context.Context, id int64) (*Order, error) {
	row := s.DB.QueryRowContext(ctx, "SELECT id, user_id, total FROM orders WHERE id = $1", id)
	var o Order
	if err := row.Scan(&o.ID, &o.UserID, &o.Total); err != nil {
		return nil, err
	}
	return &o, nil
}
