package logger

import (
	"context"
	"fmt"
)

//TODO use a real logging library
//TODO consider any headers that might be useful to log

func init() {
	//init logging library
}

func LogInfo(ctx context.Context, message string, v ...interface{}) {
	log(ctx, message, "INFO", v)
}

func LogWarn(ctx context.Context, message string, v ...interface{}) {
	log(ctx, message, "WARN", v)
}

func LogError(ctx context.Context, message string, v ...interface{}) {
	log(ctx, message, "ERROR", v)
}

func log(ctx context.Context, message string, messageType string, v ...interface{}) {
	fmt.Printf(messageType+": "+message, v...)
}
