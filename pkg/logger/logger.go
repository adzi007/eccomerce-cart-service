package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {

	// logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// // Configure zerolog logger
	// logger = zerolog.New(logFile).With().Caller().Timestamp().Logger()
	// zerolog.SetGlobalLevel(zerolog.DebugLevel) // Set global log level (optional)

	logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

}

func NewLogger() zerolog.Logger {
	return logger
}

func Trace() *zerolog.Event {
	return logger.Trace()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

func WithLevel(level zerolog.Level) *zerolog.Event {
	return logger.WithLevel(level)
}
