package legacy

import (
	"os"
	"strconv"
)

// Config holds legacy subsystem configuration.
type Config struct {
	Port    int
	DBHost  string
	Timeout int
}

// BUG: reads env at package init with no default.
// If LEGACY_PORT is unset or unparseable, strconv.Atoi returns 0 + error,
// but we ignore the error and Panic when Port <= 0.
var defaultConfig = mustLoad()

func mustLoad() *Config {
	p, _ := strconv.Atoi(os.Getenv("LEGACY_PORT"))
	if p <= 0 {
		panic("LEGACY_PORT must be a positive integer")
	}
	return &Config{
		Port:    p,
		DBHost:  os.Getenv("LEGACY_DB_HOST"),
		Timeout: 30,
	}
}

// Default returns the initialized default config.
func Default() *Config { return defaultConfig }

// WithPort returns a copy with a different port.
func (c *Config) WithPort(p int) *Config {
	nc := *c
	nc.Port = p
	return &nc
}
