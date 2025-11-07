package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

type LoggerConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	IsDev      bool
}

func NewLoggerConfig(config LoggerConfig) *zerolog.Logger {
	var writer io.Writer
	if config.IsDev {
		writer = PrettyJSONWriter{Writter: os.Stdout}
	} else {
		writer = &lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize, // megabytes
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge, // days
			Compress:   config.Compress,
			LocalTime:  true,
		}
	}
	log := zerolog.New(writer).With().Timestamp().Logger()
	return &log
}

func NewLoggerWithPath(path string, level string) *zerolog.Logger {
	config := LoggerConfig{
		Level:      level,
		Filename:   path,
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   false,
		IsDev:      GetEnv("APP_ENV") == "development",
	}
	return NewLoggerConfig(config)
}

type PrettyJSONWriter struct {
	Writter io.Writer
}

func (w PrettyJSONWriter) Write(p []byte) (n int, err error) {
	var prettyJson bytes.Buffer
	err = json.Indent(&prettyJson, p, "", "  ")
	if err != nil {
		return 0, err
	}
	return w.Writter.Write(prettyJson.Bytes())
}
