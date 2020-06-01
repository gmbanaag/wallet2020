package logger

import (
	"log"
	"path"
	"runtime"
)

//Level indicates log level
type Level uint

var logLevel uint

const (
	//LogLevelTrace = 0
	LogLevelTrace = iota
	//LogLevelDebug = 1
	LogLevelDebug
	//LogLevelInfo = 2
	LogLevelInfo
	//LogLevelError = 3
	LogLevelError
	//LogLevelFatal = 4
	LogLevelFatal
)

//Log encapsulates logging mechanism
func Log(level uint, message string) {
	if logLevel <= level {
		_, fn, line, _ := runtime.Caller(2)
		log.Printf("[%s] %s [%s:%d]", GetLogLevelString(level), message, path.Base(fn), line)
	}
}

//SetLogLevel sets service's log level
func SetLogLevel(level uint) {
	logLevel = level
}

//LogTrace wrapper for trace logs
func LogTrace(message string) {
	Log(LogLevelTrace, message)
}

//LogDebug wrapper for debug logs
func LogDebug(message string) {
	Log(LogLevelDebug, message)
}

//LogInfo wrapper for info logs
func LogInfo(message string) {
	Log(LogLevelInfo, message)
}

//LogError wrapper for error logs
func LogError(message string) {
	Log(LogLevelError, message)
}

//LogFatal wrapper for fatal logs
func LogFatal(message string) {
	Log(LogLevelFatal, message)
}

//GetLogLevelString return string to be attached in the log
func GetLogLevelString(level uint) string {

	var logLevel Level
	switch level {
	case 0:
		logLevel = LogLevelTrace
	case 1:
		logLevel = LogLevelDebug
	case 2:
		logLevel = LogLevelInfo
	case 3:
		logLevel = LogLevelError
	case 4:
		logLevel = LogLevelFatal
	default:
		logLevel = LogLevelError
	}
	return logLevel.String()
}

func (l Level) String() string {
	return [...]string{"TRACE", "DEBUG", "INFO", "ERROR", "FATAL"}[l]
}
