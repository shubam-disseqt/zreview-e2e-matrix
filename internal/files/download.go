package files

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// ServeFile streams a file from the configured storage root.
// BUG 4: path traversal — filepath.Join does not prevent "../../etc/passwd".
func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	full := filepath.Join("/var/data", name)
	f, err := os.Open(full)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()
	_, _ = io.Copy(w, f)
}
