package filesystem

import (
	"os"
	"path/filepath"
)

// Remove a Photo
func RemovePhoto(pid string) error {

	directory := "/tmp/filesystem/"
	filepath := filepath.Join(directory, pid)
	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}
