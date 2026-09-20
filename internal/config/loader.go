package config

import (
	"encoding/json"
	"os"
)

type Settings struct {
	Port     int    `json:"port"`
	LogLevel string `json:"log_level"`
}

func Load(path string) (*Settings, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var s Settings
	if err := json.NewDecoder(f).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}
