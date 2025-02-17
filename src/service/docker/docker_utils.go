package docker

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/makeopensource/leviathan/common"
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

func SaveDockerfile(filename string, contents []byte) (string, error) {
	tmpPath, err := os.CreateTemp(common.DockerFilesFolder.GetStr(), fmt.Sprintf("%s_*", filename))
	if err != nil {
		return "", err
	}

	_, err = tmpPath.Write(contents)
	if err != nil {
		return "", err
	}
	defer func(tmpPath *os.File) {
		err := tmpPath.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing temp file")
		}
	}(tmpPath)

	abs, err := filepath.Abs(tmpPath.Name())
	if err != nil {
		return "", err
	}

	return abs, nil
}

// ParseCombinedID decode combined id which should contain the machine id and container id
func ParseCombinedID(combinedId string) (string, string, error) {
	machineId, containerId, err := common.DecodeID(combinedId)
	if err != nil {
		log.Error().Err(err).Str("ID", combinedId).Msg("Could not decode ID")
		return "", "", err
	}

	return containerId, machineId, nil
}
