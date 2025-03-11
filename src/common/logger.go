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

var (
	consoleWriter = zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}
	baseLogger = log.With().Caller().Logger().Output(consoleWriter)
)

func CreateJobSubLoggerCtx(ctx context.Context, jobID string) context.Context {
	return log.Logger.With().Str(JobLogKey, jobID).Logger().WithContext(ctx)
}

func FileConsoleLogger() zerolog.Logger {
	level, err := zerolog.ParseLevel(LogLevel.GetStr())
	if err != nil {
		log.Fatal().Err(err).Msg("unable to parse log level")
	}
	log.Info().Msgf("log level is now set to %s, this can be changed by using the LEVIATHAN_LOG_LEVEL env", level)

	return baseLogger.Output(
		zerolog.MultiLevelWriter(
			GetFileLogger(LogDir.GetStr()),
			consoleWriter,
		),
	).Level(level)
}

func ConsoleLogger() zerolog.Logger {
	return baseLogger
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
