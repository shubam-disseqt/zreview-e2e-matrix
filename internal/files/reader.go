package files

import (
	"io"
	"os"
)

const UserRoot = "/var/app/users"

// ReadUserFile reads a file from the user's directory. name comes from an HTTP query param.
func ReadUserFile(userID, name string, w io.Writer) error {
	path := UserRoot + "/" + userID + "/" + name
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	return err
}
