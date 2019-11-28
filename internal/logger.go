package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Metadata map[string]string

type Logger interface {
	AddMetadata(Metadata) Logger
	Debug(string)
	DebugWithMetadata(string, Metadata)
	Info(string)
	InfoWithMetadata(string, Metadata)
	Warn(string)
	WarnWithMetadata(string, Metadata)
}

type DefaultLogger struct {
	EnableDebug bool
	logger      *log.Logger
	metadata    Metadata
}

func NewDefaultLogger(debug bool) *DefaultLogger {
	return &DefaultLogger{
		EnableDebug: debug,
		logger:      log.New(os.Stdout, "", log.LUTC),
		metadata:    make(map[string]string),
	}
}

func (l *DefaultLogger) AddMetadata(metadata Metadata) Logger {
	logger := NewDefaultLogger(l.EnableDebug)

	m := Metadata{}
	for key, value := range l.metadata {
		m[key] = value
	}
	for key, value := range metadata {
		m[key] = value
	}

	logger.metadata = m

	return logger
}

func (l *DefaultLogger) Debug(msg string) {
	if l.EnableDebug {
		l.display("DEBUG ", msg)
	}
}

func (l *DefaultLogger) DebugWithMetadata(msg string, metadata Metadata) {
	logger := l.AddMetadata(metadata)
	logger.Debug(msg)
}

func (l *DefaultLogger) Info(msg string) {
	l.display("INFO  ", msg)
}

func (l *DefaultLogger) InfoWithMetadata(msg string, metadata Metadata) {
	logger := l.AddMetadata(metadata)
	logger.Info(msg)
}

func (l *DefaultLogger) Warn(msg string) {
	l.display("WARN  ", msg)
}

func (l *DefaultLogger) WarnWithMetadata(msg string, metadata Metadata) {
	logger := l.AddMetadata(metadata)
	logger.Warn(msg)
}

func (l *DefaultLogger) display(prefix string, msg string) {
	var logmsg strings.Builder

	logmsg.WriteString(prefix)
	for key, value := range l.metadata {
		logmsg.WriteString(fmt.Sprintf("%s=%s; ", key, value))
	}
	logmsg.WriteString(msg)

	l.logger.Println(logmsg.String())
}
