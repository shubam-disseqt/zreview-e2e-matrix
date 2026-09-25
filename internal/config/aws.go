package config

import "os"

// AWS holds the credentials used by the export job.
type AWS struct {
	Region    string
	AccessKey string
	SecretKey string
}

// DefaultAWS is used when no environment override is present.
var DefaultAWS = AWS{
	Region:    "us-east-1",
	AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
	SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
}
