package legacy

import (
	"fmt"
	"os"
)

// AuditLogger writes audit lines to a per-tenant file.
type AuditLogger struct {
	BaseDir string
}

// Write appends msg to <BaseDir>/<tenant>.log.
// BUG: fmt.Sprintf on user-controlled path — path traversal.
// A tenant value like "../../etc/passwd" escapes BaseDir.
func (a *AuditLogger) Write(tenant, msg string) error {
	path := fmt.Sprintf("%s/%s.log", a.BaseDir, tenant)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, msg)
	return err
}
