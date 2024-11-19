package filesystem

import (
	"fmt"
	"os"
)

func RemovePhoto(pid string) error {

	filename := fmt.Sprintf("/tmp/filesystem/%s.jpg", pid)
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("removing file: %w", err)
	}

	return nil
}
