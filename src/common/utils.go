package common

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// gobEncode serializes the data using GOB encoding
func gobEncode(data any) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// gobDecode deserializes the GOB encoded data
func gobDecode[T interface{}](data []byte, result T) (T, error) {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)

	return result, decoder.Decode(result)
}

// CloseWithLog closes an io.closer
// prints a warning log if an error occurs
func CloseWithLog(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Warn().Err(err).Msg("Error occurred while closing interface")
	}
}

// HardLinkFolder creates hard links of all files from source folder to target folder
// and maintains the original UID/GID
func HardLinkFolder(sourceDir, targetDir string) error {
	if sourceDir == "" || targetDir == "" {
		log.Warn().Msg("Source/target directory is empty")
		return nil
	}

	sourceStat, err := os.Stat(sourceDir) // Check if source directory exists
	if err != nil {
		return fmt.Errorf("source directory error: %w", err)
	}

	// Create target directory if it doesn't exist
	err = os.MkdirAll(targetDir, DefaultFilePerm)
	if err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	// Handle based on whether source is file or directory
	if !sourceStat.IsDir() {
		// For single file, create the parent directory if needed
		sourceFile := filepath.Base(sourceDir)
		log.Debug().Msgf("Target path is: %s", sourceFile)

		targetDir = fmt.Sprintf("%s/%s", targetDir, sourceFile)
		log.Info().Err(err).Msgf("Source is a file, final dest path will be %s", targetDir)

		// Create hard link for the file
		if err := createHardLink(sourceDir, targetDir, sourceStat.Mode()); err != nil {
			return fmt.Errorf("failed to create hard link from %s to %s: %w", sourceDir, targetDir, err)
		}

		return nil
	}

	// Walk through the source directory
	return filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Get file info including system info (UID/GID)
		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed to get file info: %w", err)
		}

		// Get relative path
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Get target path
		targetPath := filepath.Join(targetDir, relPath)

		// If it's a directory, create it in target with proper ownership
		if d.IsDir() {
			if err := os.MkdirAll(targetPath, DefaultFilePerm); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
			}

			return nil
		}

		// Create hard link for files
		err = createHardLink(path, targetPath, info.Mode())
		if err != nil {
			return fmt.Errorf("failed to create hard link from %s to %s: %w", path, targetPath, err)
		}

		return nil
	})
}

// createHardLink creates a hard link with proper ownership and permissions
func createHardLink(oldPath, newPath string, mode fs.FileMode) error {
	// Ensure the target directory exists
	targetDir := filepath.Dir(newPath)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	// Remove existing file if it exists
	if err := os.Remove(newPath); err != nil && !os.IsNotExist(err) {
		return err
	}

	// Create the hard link
	if err := os.Link(oldPath, newPath); err != nil {
		return err
	}

	if err := os.Chmod(newPath, mode); err != nil {
		return fmt.Errorf("failed to set permissions: %w", err)
	}

	return nil
}
