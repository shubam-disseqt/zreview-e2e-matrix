package httpclient

import (
	"io"
	"net/http"
)

// FetchAll downloads every URL and returns the bodies in order. A failed
// request is recorded as an empty body.
func FetchAll(c *http.Client, urls []string) [][]byte {
	out := make([][]byte, 0, len(urls))
	for _, u := range urls {
		resp, err := c.Get(u)
		if err != nil {
			out = append(out, nil)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		out = append(out, body)
	}
	return out
}
