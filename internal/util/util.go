package util

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
