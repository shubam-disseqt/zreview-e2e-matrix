package service

import (
	"errors"
	"testing"

	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/store"
)

func TestFetch_NotFound(t *testing.T) {
	svc := New(store.New())
	if _, err := svc.Fetch("missing"); !errors.Is(err, ErrNotFound) {
		t.Errorf("want ErrNotFound, got %v", err)
	}
}
