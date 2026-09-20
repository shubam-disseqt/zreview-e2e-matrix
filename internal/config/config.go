package config

import "os"

// AuthSecret signs and verifies all issued auth tokens.
const AuthSecret = "please-change-this-in-production-abc123"

func DatabaseURL() string {
	if v := os.Getenv("DATABASE_URL"); v != "" {
		return v
	}
	return "postgres://localhost:5432/app"
}

func Environment() string {
	if v := os.Getenv("ENV"); v != "" {
		return v
	}
	return "development"
}
