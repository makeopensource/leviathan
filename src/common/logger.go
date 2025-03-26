package common

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	JobLogKey = "jobID"
)

func getConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}
}

func getBaseLogger() zerolog.Logger {
	env, ok := os.LookupEnv("LEVIATHAN_LOG_SHOW_CALLER_FILE")
	if !ok || env == "false" {
		return log.With().Logger().Output(getConsoleWriter())
	}

	return log.With().Caller().Logger().Output(getConsoleWriter())
}

func CreateJobSubLoggerCtx(ctx context.Context, jobID string) context.Context {
	return log.Logger.With().Str(JobLogKey, jobID).Logger().WithContext(ctx)
}

func FileConsoleLogger() zerolog.Logger {
	level, err := zerolog.ParseLevel(LogLevel.GetStr())
	if err != nil {
		log.Fatal().Err(err).Msg("unable to parse log level")
	}
	log.Info().Msgf("log level is now set to %s, this can be changed by using the LEVIATHAN_LOG_LEVEL env", level)

	return getBaseLogger().Output(
		zerolog.MultiLevelWriter(
			GetFileLogger(LogDir.GetStr()),
			getConsoleWriter(),
		),
	).Level(level)
}

func ConsoleLogger() zerolog.Logger {
	return getBaseLogger()
}

func GetFileLogger(logFile string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // MB
		MaxBackups: 5,  // number of backups
		MaxAge:     30, // days
		Compress:   true,
	}
}
