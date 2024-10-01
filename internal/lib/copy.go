/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/
package lib

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Copy(src, desc string) error {
	// Get the source information
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		// Handle directory copy
		return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Get the destination path
			destPath := filepath.Join(desc, path[len(src):])

			// If encounter dir while current source file walk through
			if info.IsDir() {
				// Create all those file and folder in the newly created destination folder path
				return os.MkdirAll(destPath, info.Mode())
			} else {
				return copyFile(path, destPath)
			}
		})
	}

	// Copy normal file
	return copyFile(src, desc)
}

func copyFile(src, desc string) error {
	fmt.Printf("Copy from %s to %s\n", src, desc)
	return nil
}
