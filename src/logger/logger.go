package logger

import (
	"fmt"
)

//TODO use a real logging library
//TODO consider any headers that might be useful to log

var (
	logFunc func(string, ...any) (n int, err error)
)

func init() {
	logFunc = fmt.Printf
}

type LogType int

const (
	Info LogType = iota
	Debug
	Warning
	Error
)

func (lt LogType) ToString() string {
	switch lt {
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		//TODO should not stop the service or panic, but should alert end users to address and fix
		return "INVALID_LOG_TYPE"
	}
}

func Log(logType LogType, message string, v ...interface{}) {
	logFunc(logType.ToString()+": "+message, v...)
}
