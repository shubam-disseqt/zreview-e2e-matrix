package config

// AWS holds the credentials used by the export job.
type AWS struct {
	Region    string
	AccessKey string
	SecretKey string
}

// DefaultAWS is used when no environment override is present.
var DefaultAWS = AWS{
	Region:    "us-east-1",
	AccessKey: "AKIA4QZX7RT2LM9WQB3F",
	SecretKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYzR9k2m4Q1x",
}
