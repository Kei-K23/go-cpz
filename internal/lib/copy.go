/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/
package lib

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/cheggaaa/pb/v3"
)

// Limit the number of concurrent copy operations
var semaphore = make(chan struct{}, 5) // Adjust the size as needed to control concurrency

func Copy(src, dest string) error {
	// Get the source information
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	if info.IsDir() {
		// Handle directory copy
		err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Get the destination path
			destPath := filepath.Join(dest, path[len(src):])

			// If encounter dir while current source file walk through
			if info.IsDir() {
				// Create all those files and folders in the newly created destination folder path
				return os.MkdirAll(destPath, info.Mode())
			}

			// If it's a file, copy it
			wg.Add(1) // Increment the wait group to track the goroutines
			copyFile(path, destPath, &wg)
			return nil // Return nil to continue walking the directory
		})
		if err != nil {
			return err
		}
	} else {
		// Copy normal file
		wg.Add(1) // Increment the wait group to track the goroutines
		copyFile(src, dest, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	return nil
}

func copyFile(src, dest string, wg *sync.WaitGroup) {
	semaphore <- struct{}{} // Acquire a spot in the semaphore to limit concurrency

	// Start the goroutine
	go func() {
		defer wg.Done()                // Decrement the wait group counter when done
		defer func() { <-semaphore }() // Release the semaphore slot

		// Open the source file
		srcFile, err := os.Open(src)
		if err != nil {
			fmt.Printf("Error opening source file: %v\n", err)
			return
		}
		defer srcFile.Close()

		// Create the destination file
		destFile, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Error creating destination file: %v\n", err)
			return
		}
		defer destFile.Close()

		// Get the metadata of the source file
		srcInfo, err := srcFile.Stat()
		if err != nil {
			fmt.Printf("Error getting source file information: %v\n", err)
			return
		}

		// Create progress indicator
		pbBar := pb.Full.Start64(srcInfo.Size())
		pbBarRdr := pbBar.NewProxyReader(srcFile)

		// Start copying the file
		_, err = io.Copy(destFile, pbBarRdr)
		if err != nil {
			fmt.Printf("Error copying file: %v\n", err)
			return
		}

		pbBar.Finish()
		fmt.Printf("Copied %s to %s\n", src, dest)
	}()
}
