package file_manager

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/internal/config"
	fu "github.com/makeopensource/leviathan/pkg/file_utils"
	"github.com/makeopensource/leviathan/pkg/logger"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
)

const (
	JobDataFolderName = "jobdata"
	DockerfileName    = "Dockerfile"
)

type FileManagerService struct{}

func NewFileManagerService() *FileManagerService {
	return &FileManagerService{}
}

type FileInfo struct {
	Reader   io.ReadCloser
	Filename string
}

// CreateTmpLabFolder lab files will be stored here
func (f *FileManagerService) CreateTmpLabFolder(dockerfile io.Reader, jobFiles ...*FileInfo) (string, error) {
	folderUUID, basePath, err := f.createBaseFolder()
	if err != nil {
		return "", err
	}

	jobDataDir := filepath.Join(basePath, JobDataFolderName)
	err = os.MkdirAll(jobDataDir, os.ModePerm)
	if err != nil {
		return "", logger.ErrLog("unable to create job data folder", err, log.Error())
	}

	if err = f.SaveFile(basePath, DockerfileName, dockerfile); err != nil {
		return "", err
	}

	for _, jobFile := range jobFiles {
		if err = f.SaveFile(jobDataDir, jobFile.Filename, jobFile.Reader); err != nil {
			return "", err
		}
	}

	return folderUUID, nil
}

// CreateSubmissionFolder submission made to grader will be places here
func (f *FileManagerService) CreateSubmissionFolder(jobFiles ...*FileInfo) (string, error) {
	folderUUID, basePath, err := f.createBaseFolder()
	if err != nil {
		return "", err
	}

	for _, jobFile := range jobFiles {
		if err = f.SaveFile(basePath, jobFile.Filename, jobFile.Reader); err != nil {
			return "", err
		}
	}

	return folderUUID, nil
}

func (f *FileManagerService) createBaseFolder() (string, string, error) {
	folderUUID, err := uuid.NewUUID()
	if err != nil {
		return "", "", logger.ErrLog("Unable to generate uuid", err, log.Error())
	}
	stringUuid := folderUUID.String()
	basePath := filepath.Join(config.TmpUploadFolder.GetStr(), stringUuid)

	err = os.Mkdir(basePath, config.DefaultFilePerm)
	if err != nil {
		return "", "", logger.ErrLog("Unable to create tmp folder", err, log.Error())
	}

	return folderUUID.String(), basePath, nil
}

func (f *FileManagerService) SaveFile(basePath string, filename string, file io.Reader) error {
	fPath := filepath.Join(basePath, filename)
	// Create the file
	dst, err := os.Create(fPath)

	if err != nil {
		return logger.ErrLog(
			"Failed to create destination file",
			err,
			log.Error(),
		)
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			log.Warn().Err(err).Msg("Error occurred while closing file")
		}
	}(dst)

	// Copy the file contents
	written, err := io.Copy(dst, file)
	if err != nil {
		return logger.ErrLog(
			"Failed to write file",
			err,
			log.Error(),
		)
	}

	log.Debug().Int64("size", written).Msgf(
		"File: %s uploaded successfully",
		filename,
	)

	return nil
}

func (f *FileManagerService) DeleteFolder(folderUuid string) {
	basePath := filepath.Join(config.TmpUploadFolder.GetStr(), folderUuid)
	if err := os.RemoveAll(basePath); err != nil {
		log.Warn().Err(err).Msgf("failed to delete tmp folder %s", folderUuid)
	}
}

func (f *FileManagerService) GetLabFilePaths(folderUuid string) (basePath string, err error) {
	basePath = filepath.Join(config.TmpUploadFolder.GetStr(), folderUuid)
	jobData := filepath.Join(basePath, JobDataFolderName)
	dockerFile := filepath.Join(basePath, DockerfileName)

	if _, err = f.checkFolder(jobData); err != nil {
		return "", err
	}
	if _, err = f.checkFolder(dockerFile); err != nil {
		return "", err
	}

	return basePath, err
}

func (f *FileManagerService) GetSubmissionPath(uuid string) (string, error) {
	path := filepath.Join(config.TmpUploadFolder.GetStr(), uuid)
	return f.checkFolder(path)
}

func (f *FileManagerService) checkFolder(path string) (jobData string, err error) {
	if !fu.FileExists(path) {
		return "", fmt.Errorf("could not find path")
	}
	return path, err
}
