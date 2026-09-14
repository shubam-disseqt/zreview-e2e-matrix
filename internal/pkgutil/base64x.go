package pkgutil

import "encoding/base64"

// EncodeURL returns the URL-safe base64 encoding of data (no padding).
func EncodeURL(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// DecodeURL decodes a URL-safe base64 string (no padding).
func DecodeURL(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
