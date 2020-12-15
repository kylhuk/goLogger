package logging

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	LogLevelError   LogLevel = 0
	LogLevelWarning LogLevel = 1 << 0
	LogLevelInfo    LogLevel = 1 << 1
	LogLevelDebug   LogLevel = 1 << 2
)

var (
	LogFull         []Log
	LogFullAsString string
)

type Log struct {
	DateTime string
	LogLevel LogLevel
	Message  string
}

//Mon Jan 2 15:04:05 -0700 MST 2006
func AddLog(level LogLevel, message string) {
	var newLog Log
	//var err error
	newLog.DateTime = time.Now().Format("2006-01-02 15:04:05.000")
	newLog.LogLevel = level
	newLog.Message = message

	LogFull = append(LogFull, newLog)
	LogFullAsString = convertToString(LogFull)
}

func convertToString(log []Log) string {
	var logString string
	var logLevel string
	for _, logEntry := range log {
		switch logEntry.LogLevel {
		case LogLevelError:
			logLevel = "ERROR"
			break
		case LogLevelWarning:
			logLevel = "WARNING"
			break
		case LogLevelInfo:
			logLevel = "INFO"
			break
		case LogLevelDebug:
			logLevel = "DEBUG"
			break
		default:
			logLevel = "N/A"
		}

		logString += fmt.Sprintf("%-25s", logEntry.DateTime) +
			fmt.Sprintf("%-9s", logLevel) +
			fmt.Sprintf("%s", logEntry.Message) + "\n"
	}

	return logString
}
