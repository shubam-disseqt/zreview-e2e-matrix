package legacy

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds legacy subsystem configuration.
type Config struct {
	Port    int
	DBHost  string
	Timeout int
}

const defaultPort = 8080

// Load returns a config from env with a safe default port.
// Fails cleanly with a descriptive error rather than panicking at init.
func Load() (*Config, error) {
	port := defaultPort
	if raw := os.Getenv("LEGACY_PORT"); raw != "" {
		p, err := strconv.Atoi(raw)
		if err != nil || p <= 0 {
			return nil, fmt.Errorf("LEGACY_PORT invalid: %q", raw)
		}
		port = p
	}
	return &Config{
		Port:    port,
		DBHost:  os.Getenv("LEGACY_DB_HOST"),
		Timeout: 30,
	}, nil
}

// WithPort returns a copy with a different port.
func (c *Config) WithPort(p int) *Config {
	nc := *c
	nc.Port = p
	return &nc
}
