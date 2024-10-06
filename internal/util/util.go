package util

import (
	"errors"
	"strings"
)

//func getFields(jsonMessage string) log.Fields {
//	fields := log.Fields{}
//	json.Unmarshal([]byte(jsonMessage), &fields)
//	return fields
//}
//
//func MultiLineResponseTrace(data string, message string) {
//	scanner := bufio.NewScanner(strings.NewReader(data))
//	for scanner.Scan() {
//		log.WithFields(getFields(scanner.Text())).Trace(message)
//	}
//}
//
//func UserHomeDir() string {
//	if runtime.GOOS == "windows" {
//		home := os.Getenv("USERPROFILE")
//		if home == "" {
//			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
//		}
//		return home
//	}
//	return os.Getenv("HOME")
//}

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
