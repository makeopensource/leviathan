package common

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const DefaultFilePerm = 0o775

// CreateTmpJobDir sets up a throwaway dir to store submission files
// you might be wondering why the '/autolab' subdir, TarDir untars it under its parent dir,
// so in container this will unpack with 'autolab' as the parent folder
// why not modify TarDir I tried and, this was easier than modifying whatever is going in that function
func CreateTmpJobDir(uuid, baseFolder string, files ...*v1.FileUpload) (string, error) {
	tmpFolder, err := os.MkdirTemp(baseFolder, uuid)
	if err != nil {
		return "", err
	}
	tmpFolder = fmt.Sprintf("%s/autolab", tmpFolder)
	err = os.MkdirAll(tmpFolder, os.ModePerm)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		err := os.WriteFile(
			fmt.Sprintf("%s/%s", tmpFolder, file.Filename),
			file.Content,
			DefaultFilePerm,
		)
		if err != nil {
			return "", err
		}
	}

	return tmpFolder, nil
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

// TarDir
// stolen from https://github.com/testcontainers/testcontainers-go/blob/f09b3af2cb985a17bd2b2eaaa5d384882ded8e28/docker.go#L633
func TarDir(src string, fileMode int64) (*bytes.Buffer, error) {
	// always pass src as absolute path
	abs, err := filepath.Abs(src)
	if err != nil {
		return &bytes.Buffer{}, fmt.Errorf("error getting absolute path: %w", err)
	}
	src = abs

	buffer := &bytes.Buffer{}

	// tar > gzip > buffer
	zr := gzip.NewWriter(buffer)
	tw := tar.NewWriter(zr)

	_, baseDir := filepath.Split(src)
	// keep the path relative to the parent directory
	index := strings.LastIndex(src, baseDir)

	// walk through every file in the folder
	err = filepath.Walk(src, func(file string, fi os.FileInfo, errFn error) error {
		if errFn != nil {
			return fmt.Errorf("error traversing the file system: %w", errFn)
		}

		// if a symlink, skip file
		if fi.Mode().Type() == os.ModeSymlink {
			log.Warn().Str("file", file).Msg("Skipping skipping symlink")
			return nil
		}

		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return fmt.Errorf("error getting file info header: %w", err)
		}

		// see https://pkg.go.dev/archive/tar#FileInfoHeader:
		// Since fs.FileInfo's Name method only returns the base name of the file it describes,
		// it may be necessary to modify Header.Name to provide the full path name of the file.
		header.Name = filepath.ToSlash(file[index:])
		header.Mode = fileMode
		now := time.Now()
		header.ModTime = now
		header.ChangeTime = now
		header.AccessTime = now

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return fmt.Errorf("error writing header: %w", err)
		}

		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return fmt.Errorf("error opening file: %w", err)
			}
			defer func(data *os.File) {
				err := data.Close()
				if err != nil {
					log.Error().Err(err).Msg("while closing file")
				}
			}(data)
			if _, err := io.Copy(tw, data); err != nil {
				return fmt.Errorf("error compressing file: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return buffer, err
	}

	// produce tar
	if err := tw.Close(); err != nil {
		return buffer, fmt.Errorf("error closing tar file: %w", err)
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return buffer, fmt.Errorf("error closing gzip file: %w", err)
	}

	return buffer, nil
}

// TarFile make input tar file from file path
// stolen from https://stackoverflow.com/a/46518557/23258902
// should not be used in CopyToContainer, which requires different file headers,
// I think I don't really know
func TarFile(filePath string) (*bytes.Reader, string, error) {
	dockerFile := filepath.Base(filePath)

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer func(tw *tar.Writer) {
		err := tw.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing tar writer")
		}
	}(tw)

	fileReader, err := os.Open(filePath)
	if err != nil {
		log.Error().Err(err).Msgf("unable to open Dockerfile")
		return nil, "", err
	}
	defer func(fileReader *os.File) {
		err := fileReader.Close()
		if err != nil {
			log.Warn().Err(err).Msg("Error while closing docker file")
		}
	}(fileReader)

	readFile, err := io.ReadAll(fileReader)
	if err != nil {
		log.Error().Err(err).Msgf("unable to read dockerfile")
		return nil, "", err
	}

	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Error().Err(err).Msgf("unable to write tar header")
		return nil, "", err
	}
	_, err = tw.Write(readFile)
	if err != nil {
		log.Error().Err(err).Msgf("unable to write tar body")
		return nil, "", err
	}

	return bytes.NewReader(buf.Bytes()), dockerFile, nil
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

// ParseCombinedID decode combined id which should contain the machine id and container id
func ParseCombinedID(combinedId string) (string, string, error) {
	machineId, containerId, err := DecodeID(combinedId)
	if err != nil {
		log.Error().Err(err).Str("ID", combinedId).Msg("Could not decode ID")
		return "", "", err
	}

	return containerId, machineId, nil
}

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

func ReadLogFile(logPath string) string {
	content, err := os.ReadFile(logPath)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to read job log file at %s", logPath)
		return ""
	}
	return string(content)
}
