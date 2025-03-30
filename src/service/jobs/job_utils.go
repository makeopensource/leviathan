package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

// CreateTmpJobDir sets up a throwaway dir to store submission files
// you might be wondering why the '/autolab' subdir, TarDir untars it under its parent dir,
// so in container this will unpack with 'autolab' as the parent folder
// why not modify TarDir I tried and, this was easier than modifying whatever is going in that function
func CreateTmpJobDir(uuid, baseFolder string) (string, error) {
	tmpFolder, err := createJobFolder(uuid, baseFolder)
	if err != nil {
		return "", err
	}

	return tmpFolder, nil
}

func createJobFolder(uuid string, baseFolder string) (string, error) {
	tmpFolder, err := os.MkdirTemp(baseFolder, uuid)
	if err != nil {
		return "", err
	}
	tmpFolder = fmt.Sprintf("%s/autolab", tmpFolder)
	if err = os.MkdirAll(tmpFolder, os.ModePerm); err != nil {
		return "", err
	}
	return tmpFolder, nil
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

func ReadLogFile(logPath string) string {
	content, err := os.ReadFile(logPath)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to read job log file at %s", logPath)
		return ""
	}
	return string(content)
}
