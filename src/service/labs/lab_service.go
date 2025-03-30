package labs

import (
	"fmt"
	. "github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	. "github.com/makeopensource/leviathan/service/file_manager"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
)

type LabService struct {
	db      *gorm.DB
	dk      *docker.DkService
	fileMan *FileManagerService
}

func NewLabService(db *gorm.DB, dk *docker.DkService, service *FileManagerService) *LabService {
	return &LabService{
		db:      db,
		dk:      dk,
		fileMan: service,
	}
}

func (service *LabService) CreateLab(lab *models.Lab, jobDirId string) (uint, error) {
	tmpDir, err := service.fileMan.GetLabFilePaths(jobDirId)
	if err != nil {
		return 0, err
	}
	defer service.fileMan.DeleteFolder(jobDirId)

	jobFolderName := fmt.Sprintf("%s_%s", lab.Name, jobDirId)
	jobDataDirPath := fmt.Sprintf("%s/%s", LabsFolder.GetStr(), jobFolderName)
	if err = os.MkdirAll(jobDataDirPath, DefaultFilePerm); err != nil {
		return 0, ErrLog(
			"unable to create directories for lab: "+lab.Name,
			err,
			log.Error(),
		)
	}

	if err = HardLinkFolder(tmpDir, jobDataDirPath); err != nil {
		return 0, ErrLog("unable to copy files to job dir", err, log.Error())
	}

	lab.DockerFilePath = filepath.Join(jobDataDirPath, DockerfileName)
	lab.JobFilesDirPath = filepath.Join(jobDataDirPath, JobDataFolderName)

	lab.ImageTag = fmt.Sprintf("%s:v1", lab.Name)
	lab.ImageTag = strings.ToLower(strings.Trim(strings.TrimSpace(lab.ImageTag), " "))

	if lab.AutolabCompatible {
		lab.JobEntryCmd = CreateTangoEntryCommand(
			WithTimeout(int(lab.JobTimeout.Seconds())),
		)
	} else {
		lab.JobEntryCmd = CreateLeviathanEntryCommand(lab.JobEntryCmd)
	}

	// final save to update paths
	db, err := service.SaveLabToDB(lab)
	if err != nil {
		return 0, err
	}

	return db.ID, nil
}

func (service *LabService) EditLab(id uint, lab *models.Lab, jobFiles string) (uint, error) {
	labData, err := service.GetLabFromDB(id)
	if err != nil {
		return 0, err
	}

	err = service.deleteLabFiles(labData)
	if err != nil {
		return 0, err
	}

	_, err = service.CreateLab(lab, jobFiles)
	if err != nil {
		return 0, err
	}

	return labData.ID, nil
}

func (service *LabService) DeleteLab(id uint) error {
	labData, err := service.GetLabFromDB(id)
	if err != nil {
		return err
	}

	err = service.deleteLabFiles(labData)
	if err != nil {
		return err
	}

	if res := service.db.Delete(&models.Lab{}, id); res.Error != nil {
		return res.Error
	}

	return nil
}

func (service *LabService) deleteLabFiles(labData *models.Lab) error {
	err := os.RemoveAll(filepath.Base(labData.DockerFilePath))
	if err != nil {
		return ErrLog(
			"unable to delete directories for lab: "+labData.Name,
			err,
			log.Error(),
		)
	}
	return nil
}

func (service *LabService) GetLabFromDB(id uint) (*models.Lab, error) {
	var lab models.Lab
	if res := service.db.Where("ID = ?", id).First(&lab); res.Error != nil {
		return nil, res.Error
	}
	return &lab, nil
}

func (service *LabService) SaveLabToDB(lab *models.Lab) (*models.Lab, error) {
	res := service.db.Save(lab)
	if res.Error != nil {
		return nil, res.Error
	}

	return lab, nil
}
