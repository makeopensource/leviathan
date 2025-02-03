package labs

import (
	"fmt"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
)

const (
	submissions = "submissions"
	grader      = "grader"
)

type LabService struct {
	db            *gorm.DB
	labFilesCache *utils.LabFilesCache
}

func NewLabService(db *gorm.DB, cache *utils.LabFilesCache) *LabService {
	return &LabService{db: db, labFilesCache: cache}
}

func (labSrv *LabService) GetLab(labName string) (*models.LabModel, error) {
	var lab models.LabModel
	res := labSrv.db.Where("labname = ?", labName).First(&lab)
	if res.Error != nil {
		log.Error().Err(res.Error)
		return &lab, fmt.Errorf("unable to query database")
	}

	return &lab, nil
}

func (labSrv *LabService) checkIfNameExists(labName string) error {
	res, err := labSrv.GetLab(labName)
	if err != nil {
		return err
	}

	if res.LabName != "" {
		return fmt.Errorf("lab already exists, use a different name")
	}

	return nil
}

func (labSrv *LabService) NewLab(lab *models.LabModel) error {
	log.Debug().Msgf("Creating course %s", lab.LabName)

	err := labSrv.checkIfNameExists(lab.LabName)
	if err != nil {
		return err
	}

	labFolder := getLabFolder(lab.LabName)
	// create grader folder
	graderFolder := fmt.Sprintf("%s/%s", labFolder, grader)
	err = createDirectoryWithOverwrite(graderFolder, false)
	if err != nil {
		return fmt.Errorf("error creating grader directory")
	}

	// write grader file
	graderFilePath := fmt.Sprintf("%s/%s", graderFolder, lab.GraderFilename)
	err = createFileWithOverwrite(graderFilePath, lab.GraderFile, false)
	if err != nil {
		return fmt.Errorf("error writing grader file")
	}

	// write makefile
	makefilePath := fmt.Sprintf("%s/%s", graderFolder, lab.MakeFilename)
	err = createFileWithOverwrite(makefilePath, lab.MakeFile, false)
	if err != nil {
		return fmt.Errorf("error writing makefile")
	}

	return nil
}

func (labSrv *LabService) EditLab(lab *models.LabModel) error {
	log.Debug().Msgf("Editing course %s", lab.LabName)

	err := labSrv.DeleteLab(lab.LabName)
	if err != nil {
		log.Error().Err(err).Msgf("error deleting lab %s", lab.LabName)
		return fmt.Errorf("error deleting lab")
	}

	err = labSrv.NewLab(lab)
	if err != nil {
		log.Error().Err(err).Msgf("error editing lab %s", lab.LabName)
		return fmt.Errorf("error editing lab")
	}

	return nil
}

func (labSrv *LabService) DeleteLab(labName string) error {
	log.Debug().Msgf("Deleting course %s", labName)

	lab, err := labSrv.GetLab(labName)
	if err != nil {
		return err
	}

	labFolder := getLabFolder(labName)

	err = os.RemoveAll(labFolder)
	if err != nil {
		log.Error().Err(err).Str("lab", labName).Msgf("failed to delete lab folder")
		return fmt.Errorf("failed to delete lab folder")
	}

	labSrv.db.Delete(&models.LabModel{}, lab.ID)

	return nil
}
