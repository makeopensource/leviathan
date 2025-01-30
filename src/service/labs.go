package service

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const submissions = "submissions"
const grader = "grader"

func NewLab(labName string, graderFilename string, graderFile []byte, makeFilename string, makeFile []byte) error {
	log.Debug().Msgf("Creating course %s", labName)

	labFolder := getLabFolder(labName)
	// create submissionFolder folder
	submissionFolder := fmt.Sprintf("%s/%s", labFolder, submissions)
	err := createDirectoryWithOverwrite(submissionFolder, true)
	if err != nil {
		return errors.New("error creating submissions directory")
	}

	// create grader folder
	graderFolder := fmt.Sprintf("%s/%s", labFolder, grader)
	err = createDirectoryWithOverwrite(graderFolder, false)
	if err != nil {
		return errors.New("error creating grader directory")
	}

	// write grader file
	graderFilePath := fmt.Sprintf("%s/%s", graderFolder, graderFilename)
	err = createFileWithOverwrite(graderFilePath, graderFile, false)
	if err != nil {
		return errors.New("error writing grader file")
	}

	// write makefile
	makefilePath := fmt.Sprintf("%s/%s", graderFolder, makeFilename)
	err = createFileWithOverwrite(makefilePath, makeFile, false)
	if err != nil {
		return errors.New("error writing makefile")
	}

	return nil
}

func EditLab(labName string, graderFilename string, graderFile []byte, makeFilename string, makeFile []byte) error {
	log.Debug().Msgf("Editing course %s", labName)

	err := DeleteLab(labName)
	if err != nil {
		log.Error().Err(err).Msgf("error deleting lab %s", labName)
		return errors.New("error deleting lab")
	}

	err = NewLab(labName, graderFilename, graderFile, makeFilename, makeFile)
	if err != nil {
		log.Error().Err(err).Msgf("error editing lab %s", labName)
		return errors.New("error editing lab")
	}

	return nil
}

func DeleteLab(labName string) error {
	log.Debug().Msgf("Deleting course %s", labName)
	labFolder := getLabFolder(labName)

	err := os.RemoveAll(labFolder)
	if err != nil {
		log.Error().Err(err).Str("lab", labName).Msgf("failed to delete lab folder")
		return errors.New("failed to delete lab folder")
	}

	return nil
}

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
			return errors.New(fmt.Sprintf("directory already exists: %s", filepath.Base(path)))
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
			return errors.New("file already exists")
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

	// Create the file
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
