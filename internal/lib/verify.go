package lib

import (
	"fmt"
	"os"
)

func Verify(src, dest string) error {
	// Get source file info
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("error accessing source file: %v", err)
	}

	// Get dest file info
	destInfo, err := os.Stat(dest)
	if err != nil {
		return fmt.Errorf("error accessing destination file: %v", err)
	}

	if srcInfo.IsDir() != destInfo.IsDir() {
		return fmt.Errorf("source and destination file type does not match")
	}

	if srcInfo.IsDir() {
		// Handle directory
	}

	return verifyFile(src, dest)
}
