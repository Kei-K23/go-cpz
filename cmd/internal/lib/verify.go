package lib

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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
		// Verify directory
		return verifyDir(src, dest)
	}

	// Verify file
	return verifyFile(src, dest)
}

// Handle when encounter directory
func verifyDir(srcDir, destDir string) error {
	err := filepath.Walk(srcDir, func(srcPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative file path for current src file
		relativePath, err := filepath.Rel(srcDir, srcPath)
		if err != nil {
			return err
		}
		fmt.Println(srcDir + "  : srcDir")
		fmt.Println(srcPath + "  : srcPath")
		fmt.Println(relativePath + " : relativePath")

		// Generate the corresponding destination relative file path from src relative file path
		destPath := filepath.Join(destDir, relativePath)
		fmt.Println(destPath + " : destPath")

		// Get the info about destination relative file
		destPathInfo, err := os.Stat(destPath)
		if err != nil {
			return fmt.Errorf("missing destination: %s", destPath)
		}

		// Check current walk through is directory, then return and end this loop
		if info.IsDir() {
			// If current walk through is directory, then destination also need to be directory
			if !destPathInfo.IsDir() {
				return fmt.Errorf("destination path %s is not a directory", destPath)
			}
			return nil // Skip the current directory
		}

		return verifyFile(srcPath, destPath)
	})

	return err
}

// Verify file with their file size and also with file content
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
