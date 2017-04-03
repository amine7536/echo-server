package common

import (
	"io"
	"log"
)

// Logging var
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// InitLogger Initialise Logger
func InitLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime)
}

// Log a message
func Log(level string, msg string) {
	switch level {
	case "INFO":
		Info.Printf(msg)
	case "TRACE":
		Trace.Printf(msg)
	case "WARNING":
		Warning.Printf(msg)
	case "ERROR":
		Error.Fatalf(msg)
	default:
		panic("Unrecognized LOG LEVEL")
	}
}
