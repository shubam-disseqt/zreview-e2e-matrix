package service

import (
	"errors"
	"fmt"

	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/store"
)

var ErrNotFound = errors.New("not found")

type Service struct {
	store *store.Store
}

func New(s *store.Store) *Service {
	return &Service{store: s}
}

func (s *Service) Fetch(id string) (store.Item, error) {
	it, ok := s.store.Get(id)
	if !ok {
		return store.Item{}, fmt.Errorf("fetch %q: %w", id, ErrNotFound)
	}
	return it, nil
}
