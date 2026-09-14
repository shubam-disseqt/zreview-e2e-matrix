package pkgutil

import "encoding/json"

// MustMarshal returns the JSON encoding of v, panicking on error.
// Use only with values known to be marshalable (config, tests, fixtures).
func MustMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic("pkgutil: MustMarshal: " + err.Error())
	}
	return b
}
