package v1

import (
	"encoding/json"
	com "github.com/makeopensource/leviathan/common"
	fm "github.com/makeopensource/leviathan/service/file_manager"
	"github.com/rs/zerolog/log"
	"mime/multipart"
	"net/http"
)

const (
	maxMemory          = 32 << 20 // 32 MB for multipart forms
	SubmissionFilesKey = "submissionFiles"
	DockerFileKey      = "dockerfile"
	LabFilesKey        = "labFiles"
)

type FileManagerHandler struct {
	BasePath             string
	UploadLabPath        string
	UploadSubmissionPath string
	service              fm.FileManagerService
}

func NewFileManagerHandler(basePath string) *FileManagerHandler {
	return &FileManagerHandler{
		BasePath:             basePath,
		UploadLabPath:        basePath + "/upload/lab",
		UploadSubmissionPath: basePath + "/upload/submission",
		service:              fm.FileManagerService{},
	}
}

func (f *FileManagerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.URL.String() == f.UploadLabPath {
			f.UploadLabData(w, r)
			return
		} else if r.URL.String() == f.UploadSubmissionPath {
			f.UploadSubmissionData(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (f *FileManagerHandler) UploadLabData(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form with reasonable memory limit
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	dockerFile, _, err := r.FormFile(DockerFileKey)
	if err != nil {
		http.Error(
			w,
			com.ErrLog("Failed to get dockerfile in form", err, log.Error()).Error(),
			http.StatusBadRequest,
		)
		return
	}
	defer com.CloseFile(dockerFile)

	jobFiles, ok := r.MultipartForm.File[LabFilesKey]
	if !ok || len(jobFiles) == 0 {
		http.Error(w, "No jobFiles in form", http.StatusBadRequest)
		return
	}

	fileInfos, err := mapToFileInfo(jobFiles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(files []*fm.FileInfo) {
		for _, file := range files {
			com.CloseFile(file.Reader)
		}
	}(fileInfos)

	folderID, err := f.service.CreateTmpLabFolder(dockerFile, fileInfos...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, folderID)
}

func (f *FileManagerHandler) UploadSubmissionData(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form with reasonable memory limit
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	jobFiles, ok := r.MultipartForm.File[SubmissionFilesKey]
	if !ok || len(jobFiles) == 0 {
		http.Error(w, "No submission files found in form", http.StatusBadRequest)
		return
	}

	fileInfos, err := mapToFileInfo(jobFiles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(files []*fm.FileInfo) {
		for _, file := range files {
			com.CloseFile(file.Reader)
		}
	}(fileInfos)

	folderID, err := f.service.CreateSubmissionFolder(fileInfos...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, folderID)
}

func sendResponse(w http.ResponseWriter, folderID string) {
	jsonData, err := toJson(folderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(
			w,
			com.ErrLog("Failed to write response", err, log.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}
}

func toJson(folderID string) ([]byte, error) {
	resultMap := map[string]string{
		"folderId": folderID,
	}
	jsonData, err := json.Marshal(resultMap)
	if err != nil {
		return nil, com.ErrLog("Failed to marshal json", err, log.Error())
	}
	return jsonData, nil
}

func mapToFileInfo(jobFiles []*multipart.FileHeader) ([]*fm.FileInfo, error) {
	var fileInfos []*fm.FileInfo
	for _, jobFile := range jobFiles {
		file, err := jobFile.Open()
		if err != nil {
			return fileInfos, com.ErrLog("unable to open file: "+err.Error(), err, log.Error())
		}
		fileInfos = append(fileInfos, &fm.FileInfo{
			Reader:   file,
			Filename: jobFile.Filename,
		})
	}
	return fileInfos, nil
}
