package utils

import (
	"archive/tar"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/pkg/archive"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const DefaultFilePerm = 0o775

func EncodeID(id1 string, id2 string) string {
	return id1 + "#" + id2
}

func DecodeID(combinedId string) (string, string, error) {
	strs := strings.Split(combinedId, "#")
	if len(strs) != 2 {
		return "", "", errors.New("could not decode ID")
	}
	return strs[0], strs[1], nil
}

// ArchiveJobData creates a tar.gz archive from the provided file map.3
func ArchiveJobData(files map[string][]byte) (string, error) {
	tmpFolder, err := os.MkdirTemp(SubmissionTarFolder.GetStr(), "submission_*")
	tmpFile, err := os.Create(fmt.Sprintf("%s/%s", tmpFolder, "grader.tar.gz"))
	if err != nil {
		return "", fmt.Errorf("creating temp file: %w", err)
	}
	defer func(tmpFile *os.File) {
		err := tmpFile.Close()
		if err != nil {
			log.Error().Err(err).Msg("while closing temp file")
		}
	}(tmpFile)

	//gz := gzip.NewWriter(tmpFile)
	tw := tar.NewWriter(tmpFile)

	for name, content := range files {
		header := &tar.Header{
			Name: name,
			Mode: DefaultFilePerm,
			Size: int64(len(content)),
		}
		if err := tw.WriteHeader(header); err != nil {
			return "", fmt.Errorf("writing header: %w", err)
		}
		if _, err := tw.Write(content); err != nil {
			return "", fmt.Errorf("writing content: %w", err)
		}
	}

	if err := tw.Close(); err != nil {
		return "", fmt.Errorf("closing tar writer: %w", err)
	}
	//if err := gz.Close(); err != nil {
	//	return "", fmt.Errorf("closing gzip writer: %w", err)
	//}

	return filepath.Abs(tmpFile.Name())
}

func ArchiveFiles(filePath string) (io.ReadCloser, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		log.Error().Err(err).Msgf("failed to stat path %s", filePath)
		return nil, err
	}

	archiveData, err := archive.Tar(filePath, archive.Gzip)
	if err != nil {
		log.Error().Err(err).Msgf("failed to archive %s", filePath)
		return nil, err
	}

	return archiveData, nil
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

// ReadFileBytes reads a file at the given filepath and returns its content as a byte slice.
// It handles errors and returns an error if the file cannot be read.
func ReadFileBytes(filepath string) ([]byte, error) {
	// Check if the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", filepath)
	}

	// Open the file for reading
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err) // Wrap the error
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error().Err(err).Msgf("error closing file: %s", filepath)
		}
	}(file) // Important: Close the file when done

	// Read all bytes from the file
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err) // Wrap the error
	}

	return fileBytes, nil
}
