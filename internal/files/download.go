package files

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const storageRoot = "/var/data"

// ServeFile streams a file from the configured storage root, rejecting any
// name that would escape the root via ".." or absolute paths.
func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" || strings.Contains(name, "..") || strings.HasPrefix(name, "/") {
		http.Error(w, "bad name", http.StatusBadRequest)
		return
	}
	full := filepath.Join(storageRoot, name)
	// Belt-and-braces: ensure the cleaned path stays inside the root.
	if !strings.HasPrefix(full, storageRoot+string(os.PathSeparator)) {
		http.Error(w, "bad name", http.StatusBadRequest)
		return
	}
	f, err := os.Open(full)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()
	_, _ = io.Copy(w, f)
}
