package dockerclient

import (
	"archive/tar"
	"bytes"
	"github.com/makeopensource/leviathan/utils"
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

func saveDockerfile(fullPath string, contents []byte) error {

	log.Debug().Str("filename", fullPath).Msgf("Recivied new container request")

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		log.Error().Err(err).Msgf("Failed to create file and folder at %s", fullPath)
		return err
	}

	err := os.WriteFile(fullPath, contents, 0644)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to write contents to file")
		return err
	}

	return nil
}

// parseCombinedID decode combined id which should contain the machine id and container id
func parseCombinedID(combinedId string) (string, string, error) {
	machineId, containerId, err := utils.DecodeID(combinedId)
	if err != nil {
		log.Error().Err(err).Str("ID", combinedId).Msg("Could not decode ID")
		return "", "", err
	}

	return containerId, machineId, nil
}
