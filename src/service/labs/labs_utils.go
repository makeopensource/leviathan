package labs

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func getLabFolder(labName string) string {
	appdataFolder := filepath.Dir(viper.ConfigFileUsed())
	labFolder := fmt.Sprintf("%s/labs/%s", appdataFolder, labName)
	return labFolder
}

// createDirectoryWithOverwrite creates a directory at the given path.
// If the directory exists and overwrite is true, it deletes the existing directory and creates a new one.
// If the directory exists and overwrite is false, it returns an error.
func createDirectoryWithOverwrite(path string, overwrite bool) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if !overwrite {
			return fmt.Errorf("directory already exists: %s", filepath.Base(path))
		}
		if err := os.RemoveAll(path); err != nil {
			log.Error().Err(err).Str("path", path).Msgf("failed to remove directory")
			return err
		}
	}
	return os.MkdirAll(path, 0755)
}

// createFileWithOverwrite creates a file at the given path.
// If the file exists and overwrite is true, it deletes the existing file and creates a new one.
// If the file exists and overwrite is false, it returns an error.
func createFileWithOverwrite(path string, contents []byte, overwrite bool) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if !overwrite {
			return fmt.Errorf("file already exists")
		}
		if err := os.Remove(path); err != nil {
			log.Error().Err(err).Str("path", path).Msgf("failed to remove directory")
			return err
		}
	}

	// Ensure the directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Error().Err(err).Str("dir", dir).Msgf("failed to create dir")
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		log.Error().Err(err).Str("path", path).Msgf("failed to create file")
		return err
	}

	err = os.WriteFile(file.Name(), contents, 0644)
	if err != nil {
		log.Error().Err(err).Str("path", path).Msgf("failed to write file")
		return err
	}

	err = file.Close()
	if err != nil {
		log.Error().Err(err).Str("path", path).Msgf("failed to close file")
		return err
	}

	return nil
}
