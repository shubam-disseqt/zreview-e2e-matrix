package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID    int64
	Email string
}

type UserStore struct{ DB *sql.DB }

func (s *UserStore) Get(ctx context.Context, id int64) (*User, error) {
	row := s.DB.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	var u User
	if err := row.Scan(&u.ID, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}
