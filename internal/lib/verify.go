package lib

import (
	"crypto/md5"
	"fmt"
	"io"
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

func verifyFile(srcFile, destFile string) error {
	// Get source file info
	srcInfo, err := os.Stat(srcFile)
	if err != nil {
		return fmt.Errorf("error accessing source file: %v", err)
	}

	// Get dest file info
	destInfo, err := os.Stat(destFile)
	if err != nil {
		return fmt.Errorf("error accessing destination file: %v", err)
	}

	// Compare the src and dest file size
	if srcInfo.Size() != destInfo.Size() {
		return fmt.Errorf("file size does not match %s and %s", srcFile, destFile)
	}

	srcHash, err := hashFile(srcFile)
	if err != nil {
		return fmt.Errorf("error hashing source file: %v", err)
	}

	destHash, err := hashFile(destFile)
	if err != nil {
		return fmt.Errorf("error hashing destination file: %v", err)
	}

	// Compare with hash values
	if srcHash != destHash {
		return fmt.Errorf("file content does not match %s and %s", srcFile, destFile)
	}

	// Nill means src and dest file are same and verify
	return nil
}

// Hash the file content to compare between file contents hash value
func hashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()

	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	// Return hash value into hex value
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
