package main

import (
	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/service"
	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/store"
)

func newService() *service.Service {
	return service.New(store.New())
}
