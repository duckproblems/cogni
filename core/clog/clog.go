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

var Level LogLevel = WARN

func logMessage(logLevel LogLevel, message string) {
	if logLevel != FATAL && logLevel < Level {
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	prefix := levelPrefix[logLevel]

	fmt.Printf("%s : %s\t> %s\n", currentTime, prefix, message)

	if logLevel == FATAL {
		os.Exit(1)
	}
}

func Debug(message string) {
	logMessage(DEBUG, message)
}

func Info(message string) {
	logMessage(INFO, message)
}

func Warn(message string) {
	logMessage(WARN, message)
}

func Error(message string) {
	logMessage(ERROR, message)
}

func Fatal(message string) {
	logMessage(FATAL, message)
}

func Debugf(format string, args ...any) {
	logMessage(DEBUG, fmt.Sprintf(format, args...))
}

func Infof(format string, args ...any) {
	logMessage(INFO, fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...any) {
	logMessage(WARN, fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...any) {
	logMessage(ERROR, fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...any) {
	logMessage(FATAL, fmt.Sprintf(format, args...))
}
