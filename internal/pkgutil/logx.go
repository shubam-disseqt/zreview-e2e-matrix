package pkgutil

import (
	"log/slog"
	"os"
)

// Logger is a thin wrapper around slog.Logger to keep call sites terse.
type Logger struct {
	sl *slog.Logger
}

// NewLogger returns a Logger writing JSON to stderr at the given level.
func NewLogger(level slog.Level) *Logger {
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	return &Logger{sl: slog.New(h)}
}

// Info logs at info level with key/value attributes.
func (l *Logger) Info(msg string, args ...any) { l.sl.Info(msg, args...) }

// Error logs at error level with key/value attributes.
func (l *Logger) Error(msg string, args ...any) { l.sl.Error(msg, args...) }
