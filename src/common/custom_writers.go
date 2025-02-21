package common

// LogWriter logger for writer
type LogWriter struct {
	LoggerFunc func(string)
}

func (z *LogWriter) Write(p []byte) (n int, err error) {
	z.LoggerFunc(string(p))
	return len(p), nil
}
