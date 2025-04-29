package labs

import (
	"fmt"
	"github.com/makeopensource/leviathan/internal/config"
	"github.com/makeopensource/leviathan/internal/docker"
	fm "github.com/makeopensource/leviathan/internal/file_manager"
	"github.com/makeopensource/leviathan/pkg/file_utils"
	"github.com/makeopensource/leviathan/pkg/logger"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

type LabService struct {
	db      LabStore
	dk      *docker.DkService
	fileMan *fm.FileManagerService
}

func NewLabService(db LabStore, dk *docker.DkService, service *fm.FileManagerService) *LabService {
	return &LabService{
		db:      db,
		dk:      dk,
		fileMan: service,
	}
}

func (service *LabService) CreateLab(lab *Lab, jobDirId string) (uint, error) {
	tmpDir, err := service.fileMan.GetLabFilePaths(jobDirId)
	if err != nil {
		return 0, err
	}
	defer service.fileMan.DeleteFolder(jobDirId)

	jobFolderName := fmt.Sprintf("%s_%s", lab.Name, jobDirId)
	jobDataDirPath := fmt.Sprintf("%s/%s", config.LabsFolder.GetStr(), jobFolderName)
	if err = os.MkdirAll(jobDataDirPath, config.DefaultFilePerm); err != nil {
		return 0, logger.ErrLog(
			"unable to create directories for lab: "+lab.Name,
			err,
			log.Error(),
		)
	}

	if err = file_utils.HardLinkFolder(tmpDir, jobDataDirPath); err != nil {
		return 0, logger.ErrLog("unable to copy files to job dir", err, log.Error())
	}

	lab.DockerFilePath = filepath.Join(jobDataDirPath, fm.DockerfileName)
	lab.JobFilesDirPath = filepath.Join(jobDataDirPath, fm.JobDataFolderName)

	lab.ImageTag = fmt.Sprintf("%s:v1", lab.Name)
	lab.ImageTag = strings.ToLower(strings.Trim(strings.TrimSpace(lab.ImageTag), " "))

	if lab.AutolabCompatible {
		lab.JobEntryCmd = createTangoEntryCommand(
			WithTimeout(int(lab.JobTimeout.Seconds())),
		)
	} else {
		lab.JobEntryCmd = createLeviathanEntryCommand(lab.JobEntryCmd)
	}

	// final save to update paths
	if err = service.db.CreateLab(lab); err != nil {
		return 0, err
	}

	return lab.ID, nil
}

func (service *LabService) EditLab(id uint, lab *Lab, jobFiles string) (uint, error) {
	labData, err := service.db.GetLab(id)
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
	labData, err := service.db.GetLab(id)
	if err != nil {
		return err
	}

	err = service.deleteLabFiles(labData)
	if err != nil {
		return err
	}

	if err = service.db.DeleteLab(id); err != nil {
		return err
	}

	return nil
}

func (service *LabService) deleteLabFiles(labData *Lab) error {
	err := os.RemoveAll(filepath.Base(labData.DockerFilePath))
	if err != nil {
		return logger.ErrLog(
			"unable to delete directories for lab: "+labData.Name,
			err,
			log.Error(),
		)
	}
	return nil
}

func (service *LabService) GetLabFromDB(id uint) (*Lab, error) {
	return service.db.GetLab(id)
}
