package pkgutil

import (
	"errors"
	"fmt"
)

// Wrap wraps err with msg. Returns nil if err is nil.
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", msg, err)
}

// Is reports whether any error in err's chain matches target.
func Is(err, target error) bool {
	return errors.Is(err, target)
}
