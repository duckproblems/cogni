package clog

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota + 1
	INFO
	WARN
	ERROR
	FATAL
)

var levelPrefix = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

type Logger struct {
	level LogLevel
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

func (l *Logger) logMessage(logLevel LogLevel, message string) {
	if logLevel != FATAL && logLevel < l.level {
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	prefix := levelPrefix[logLevel]

    fmt.Printf("%s : %s\t> %s\n", currentTime, prefix, message)

	if logLevel == FATAL {
		os.Exit(1)
	}
}

func (l *Logger) Debug(message string) {
	l.logMessage(DEBUG, message)
}

func (l *Logger) Info(message string) {
	l.logMessage(INFO, message)
}

func (l *Logger) Warn(message string) {
	l.logMessage(WARN, message)
}

func (l *Logger) Error(message string) {
	l.logMessage(ERROR, message)
}

func (l *Logger) Fatal(message string) {
	l.logMessage(FATAL, message)
}
