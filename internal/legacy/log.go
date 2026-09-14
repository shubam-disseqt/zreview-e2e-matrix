package legacy

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// AuditLogger writes audit lines to a per-tenant file.
type AuditLogger struct {
	BaseDir string
}

// ErrInvalidTenant is returned when a tenant identifier is unsafe.
var ErrInvalidTenant = errors.New("invalid tenant identifier")

// Write appends msg to <BaseDir>/<tenant>.log.
// Tenant is validated to prevent path traversal.
func (a *AuditLogger) Write(tenant, msg string) error {
	if tenant == "" || strings.ContainsAny(tenant, `/\`) || strings.Contains(tenant, "..") {
		return ErrInvalidTenant
	}
	base, err := filepath.Abs(a.BaseDir)
	if err != nil {
		return err
	}
	path := filepath.Join(base, tenant+".log")
	if !strings.HasPrefix(path, base+string(os.PathSeparator)) {
		return ErrInvalidTenant
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, msg)
	return err
}
