package logger

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		logType  LogType
		message  string
		expected string
	}{
		{logType: Info, message: "This is an info message", expected: "INFO: This is an info message"},
		{logType: Debug, message: "This is a debug message", expected: "DEBUG: This is a debug message"},
		{logType: Warning, message: "This is a warning message", expected: "WARNING: This is a warning message"},
		{logType: Error, message: "This is an error message", expected: "ERROR: This is an error message"},
		{logType: LogType(999), message: "This is an invalid log type message", expected: "INVALID_LOG_TYPE: This is an invalid log type message"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			var buf bytes.Buffer
			logFunc = func(format string, args ...interface{}) (n int, err error) {
				return fmt.Fprintf(&buf, format, args...)
			}
			defer func() { logFunc = fmt.Printf }()

			Log(tt.logType, tt.message)

			if buf.String() != tt.expected {
				t.Errorf("expected %q but got %q", tt.expected, buf.String())
			}
		})
	}
}
