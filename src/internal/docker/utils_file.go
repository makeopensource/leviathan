package docker

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// tarDir
// stolen from https://github.com/testcontainers/testcontainers-go/blob/f09b3af2cb985a17bd2b2eaaa5d384882ded8e28/docker.go#L633
func tarDir(src string, fileMode int64) (*bytes.Buffer, error) {
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

// tarFile make input tar file from file path
// stolen from https://stackoverflow.com/a/46518557/23258902
// should not be used in CopyToContainer, which requires different file headers,
// I think I don't really know
func tarFile(filePath string) (*bytes.Reader, string, error) {
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
