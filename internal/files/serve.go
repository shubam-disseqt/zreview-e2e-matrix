package files

import (
	"net/http"
	"os"
	"path/filepath"
)

// Root is the directory user downloads are served from.
var Root = "/srv/uploads"

// Serve writes the requested file to the response.
func Serve(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	p := filepath.Join(Root, name)
	data, err := os.ReadFile(p)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	_, _ = w.Write(data)
}
