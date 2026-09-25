package httpclient

import (
	"crypto/tls"
	"net/http"
	"time"
)

// New returns the HTTP client used for all outbound calls.
func New() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr, Timeout: 10 * time.Second}
}
