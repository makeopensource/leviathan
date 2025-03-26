package labs

import (
	"fmt"
	com "github.com/makeopensource/leviathan/common"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

type LabService struct {
	db *gorm.DB
	dk *docker.DkService
}

func NewLabService(db *gorm.DB, dk *docker.DkService) *LabService {
	return &LabService{
		db: db,
		dk: dk,
	}
}

func (service *LabService) CreateLab(lab *models.Lab, dockerFile *v1.FileUpload, jobFiles []*v1.FileUpload) (uint, error) {
	// first save to check lab name uniqueness
	lab, err := service.SaveLabToDB(lab)
	if err != nil {
		return 0, err
	}

	basePath := fmt.Sprintf("%s/%s", com.LabsFolder.GetStr(), strconv.Itoa(int(lab.ID)))
	jobDataDirPath := basePath + "/jobData"

	if err = os.MkdirAll(jobDataDirPath, com.DefaultFilePerm); err != nil {
		log.Error().Err(err).Msgf("unable to create directories for lab: %s", lab.Name)
		return 0, fmt.Errorf("unable to create directories for lab: %s", lab.Name)
	}

	for _, file := range jobFiles {
		path := fmt.Sprintf("%s/%s", jobDataDirPath, file.Filename)
		if err := writeFile(path, file.Content); err != nil {
			return 0, err
		}
	}

	lab.ImageTag = fmt.Sprintf("%s:v1", lab.Name)
	lab.ImageTag = strings.ToLower(strings.Trim(strings.TrimSpace(lab.ImageTag), " "))

	dockerFile.Filename = fmt.Sprintf("%s_%s", lab.Name, dockerFile.Filename)
	dockerFilePath := fmt.Sprintf("%s/%s", basePath, dockerFile.Filename)

	if err = writeFile(dockerFilePath, dockerFile.Content); err != nil {
		return 0, err
	}

	lab.DockerFilePath = dockerFilePath
	lab.JobFilesDirPath = jobDataDirPath

	// final save to update paths
	db, err := service.SaveLabToDB(lab)
	if err != nil {
		return 0, err
	}

	return db.ID, nil
}

func (service *LabService) EditLab(id uint, lab *models.Lab) error {
	panic("implement me")
}

func (service *LabService) DeleteLab(id uint) error {
	if res := service.db.Delete(&models.Lab{}, id); res.Error != nil {
		return res.Error
	}
	return nil
}

func (service *LabService) GetLabFromDB(id uint) (*models.Lab, error) {
	var lab models.Lab
	if res := service.db.First(&lab).Where("ID = ?", id); res.Error != nil {
		return nil, res.Error
	}
	return &lab, nil
}

func (service *LabService) SaveLabToDB(lab *models.Lab) (*models.Lab, error) {
	if res := service.db.Save(lab); res.Error != nil {
		return nil, res.Error
	}
	return lab, nil
}

func writeFile(path string, fileData []byte) error {
	return os.WriteFile(
		path,
		fileData,
		com.DefaultFilePerm,
	)
}
