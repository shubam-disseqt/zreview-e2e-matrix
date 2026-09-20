package store

import (
	"database/sql"
	"fmt"
)

type ProductStore struct{ DB *sql.DB }

func (s *ProductStore) FindByName(query string) ([]Product, error) {
	q := fmt.Sprintf("SELECT id, name, price FROM products WHERE name LIKE '%%%s%%'", query)
	rows, err := s.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}

type Product struct {
	ID    int64
	Name  string
	Price float64
}
