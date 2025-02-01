package labs

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
)

const (
	submissions = "submissions"
	grader      = "grader"
)

type LabService struct {
	db *gorm.DB
}

func NewLabService(db *gorm.DB) *LabService {
	return &LabService{db: db}
}

func (labSrv *LabService) NewLab(labName string, graderFilename string, graderFile []byte, makeFilename string, makeFile []byte) error {
	log.Debug().Msgf("Creating course %s", labName)

	labFolder := getLabFolder(labName)
	// create submissionFolder folder
	submissionFolder := fmt.Sprintf("%s/%s", labFolder, submissions)
	err := createDirectoryWithOverwrite(submissionFolder, true)
	if err != nil {
		return fmt.Errorf("error creating submissions directory")
	}

	// create grader folder
	graderFolder := fmt.Sprintf("%s/%s", labFolder, grader)
	err = createDirectoryWithOverwrite(graderFolder, false)
	if err != nil {
		return fmt.Errorf("error creating grader directory")
	}

	// write grader file
	graderFilePath := fmt.Sprintf("%s/%s", graderFolder, graderFilename)
	err = createFileWithOverwrite(graderFilePath, graderFile, false)
	if err != nil {
		return fmt.Errorf("error writing grader file")
	}

	// write makefile
	makefilePath := fmt.Sprintf("%s/%s", graderFolder, makeFilename)
	err = createFileWithOverwrite(makefilePath, makeFile, false)
	if err != nil {
		return fmt.Errorf("error writing makefile")
	}

	return nil
}

func (labSrv *LabService) EditLab(labName string, graderFilename string, graderFile []byte, makeFilename string, makeFile []byte) error {
	log.Debug().Msgf("Editing course %s", labName)

	err := labSrv.DeleteLab(labName)
	if err != nil {
		log.Error().Err(err).Msgf("error deleting lab %s", labName)
		return fmt.Errorf("error deleting lab")
	}

	err = labSrv.NewLab(labName, graderFilename, graderFile, makeFilename, makeFile)
	if err != nil {
		log.Error().Err(err).Msgf("error editing lab %s", labName)
		return fmt.Errorf("error editing lab")
	}

	return nil
}

func (labSrv *LabService) DeleteLab(labName string) error {
	log.Debug().Msgf("Deleting course %s", labName)
	labFolder := getLabFolder(labName)

	err := os.RemoveAll(labFolder)
	if err != nil {
		log.Error().Err(err).Str("lab", labName).Msgf("failed to delete lab folder")
		return fmt.Errorf("failed to delete lab folder")
	}

	return nil
}
