package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/pkg/archive"
	"github.com/rs/zerolog/log"
	"io"
	"os"
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

// ArchiveJobData creates a tar.gz archive from the provided file map.
func ArchiveJobData(files map[string][]byte) (io.ReadCloser, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)

	var filePerm int64 = 0600

	// Add each file to the archive
	for name, content := range files {
		header := &tar.Header{
			Name: name,
			Mode: filePerm,
			Size: int64(len(content)),
		}
		if err := tw.WriteHeader(header); err != nil {
			return nil, err
		}
		if _, err := tw.Write(content); err != nil {
			return nil, err
		}
	}
	if err := tw.Close(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}

	// we create a new reader so that
	// we can close and cleanup the above tar/gzip writers
	return io.NopCloser(bytes.NewReader(buf.Bytes())), nil
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
