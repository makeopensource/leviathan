package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/makeopensource/leviathan/common"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/rs/zerolog/log"
	"io/fs"
	"os"
	"path/filepath"
)

// HardLinkFolder creates hard links of all files from source folder to target folder
// and maintains the original UID/GID
func HardLinkFolder(sourceDir, targetDir string) error {
	// Check if source directory exists
	sourceStat, err := os.Stat(sourceDir)
	if err != nil {
		return fmt.Errorf("source directory error: %w", err)
	}

	// Create target directory if it doesn't exist
	err = os.MkdirAll(targetDir, common.DefaultFilePerm)
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
			if err := os.MkdirAll(targetPath, common.DefaultFilePerm); err != nil {
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

// CreateTmpJobDir sets up a throwaway dir to store submission files
// you might be wondering why the '/autolab' subdir, TarDir untars it under its parent dir,
// so in container this will unpack with 'autolab' as the parent folder
// why not modify TarDir I tried and, this was easier than modifying whatever is going in that function
func CreateTmpJobDir(uuid, baseFolder string, files ...*v1.FileUpload) (string, error) {
	tmpFolder, err := createJobFolder(uuid, baseFolder)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if err := os.WriteFile(
			fmt.Sprintf("%s/%s", tmpFolder, file.Filename),
			file.Content,
			common.DefaultFilePerm,
		); err != nil {
			return "", err
		}
	}

	return tmpFolder, nil
}

func createJobFolder(uuid string, baseFolder string) (string, error) {
	tmpFolder, err := os.MkdirTemp(baseFolder, uuid)
	if err != nil {
		return "", err
	}
	tmpFolder = fmt.Sprintf("%s/autolab", tmpFolder)
	if err = os.MkdirAll(tmpFolder, os.ModePerm); err != nil {
		return "", err
	}
	return tmpFolder, nil
}

func GetLastLine(file *os.File) (string, error) {
	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	var lastLine string
	buf := make([]byte, 1)
	offset := stat.Size() - 1

	for {
		if offset < 0 {
			break
		}

		_, err := file.Seek(offset, 0)
		if err != nil {
			return "", err
		}

		_, err = file.Read(buf)
		if err != nil {
			return "", err
		}

		if buf[0] == '\n' && lastLine != "" {
			break
		}

		lastLine = string(buf) + lastLine
		offset--
	}

	if lastLine == "" {
		return "", fmt.Errorf("last line is empty")
	}

	return lastLine, nil
}

func IsValidJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

func ReadLogFile(logPath string) string {
	content, err := os.ReadFile(logPath)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to read job log file at %s", logPath)
		return ""
	}
	return string(content)
}
