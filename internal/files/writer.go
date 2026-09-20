package files

import (
	"io"
	"os"
	"path/filepath"
)

// WriteBlob writes data to a blob store. The key is an internally-generated
// UUID (never user input) so path composition is safe here.
func WriteBlob(root, key string, data io.Reader) error {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(root, key))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, data)
	return err
}
