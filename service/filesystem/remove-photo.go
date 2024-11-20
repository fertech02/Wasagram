package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

const uploadDirectory = "/tmp/filesystem/"

func RemovePhoto(photoID string) error {

	filename := filepath.Join(uploadDirectory, fmt.Sprintf("%s%s", photoID, ".jpg"))
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("removing file: %w", err)
	}

	return nil
}
