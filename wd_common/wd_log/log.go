package wd_log

import (
	"github.com/rs/zerolog"
	"os"
	"strconv"
)

func New(path, fileName string, logLevel zerolog.Level) zerolog.Logger {
	console := zerolog.ConsoleWriter{Out: os.Stderr}
	multi := zerolog.MultiLevelWriter(console, NewFileWriter(path, fileName))
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return file + ":" + strconv.Itoa(line)
	}
	return zerolog.New(multi).Level(logLevel).With().Timestamp().Caller().Logger()
}
