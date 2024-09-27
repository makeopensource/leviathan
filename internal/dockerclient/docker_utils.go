package dockerclient

import (
	"archive/tar"
	"bytes"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
)

// ConvertToTar make input tar file from dockerfile path
// courtesy of https://stackoverflow.com/a/46518557/23258902
func ConvertToTar(dockerFilePath string) (*bytes.Reader, string) {
	dockerFile := filepath.Base(dockerFilePath)
	log.Debug().Msgf("Docker file: %s: from path %s", dockerFile, dockerFilePath)

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer func(tw *tar.Writer) {
		err := tw.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing tar writer")
		}
	}(tw)

	dockerFileReader, err := os.Open(dockerFilePath)
	if err != nil {
		log.Error().Err(err).Msgf("unable to open Dockerfile")
		return nil, ""
	}
	readDockerFile, err := io.ReadAll(dockerFileReader)
	if err != nil {
		log.Error().Err(err).Msgf("unable to read dockerfile")
		return nil, ""
	}

	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Error().Err(err).Msgf("unable to write tar header")
		return nil, ""
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Error().Err(err).Msgf("unable to write tar body")
		return nil, ""
	}

	return bytes.NewReader(buf.Bytes()), dockerFile
}
