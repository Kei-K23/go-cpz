package lib

import (
	"fmt"
	"os"
)

func Copy(src, desc string) error {
	// Get the source information
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		// Handle directory copy
		return copyFile(src, desc)
	} else {
		// Handle file copy
		return copyFile(src, desc)
	}
}

func copyFile(src, desc string) error {
	fmt.Printf("Copy from %s to %s\n", src, desc)
	return nil
}
