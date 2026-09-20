package files

import (
	"io"
	"os"
)

const UploadRoot = "/var/app/uploads"

// SaveUpload writes an uploaded file to the per-user upload directory.
// filename is taken from the multipart form field's Filename.
func SaveUpload(userID, filename string, body io.Reader) error {
	path := UploadRoot + "/" + userID + "/" + filename
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, body)
	return err
}
