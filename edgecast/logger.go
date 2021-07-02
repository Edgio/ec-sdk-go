// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
package edgecast

import (
	"log"
	"os"
)

// Logger defines the the basic implementation that meets the SDK's logging needs
type Logger interface {
	// LogDebug writes debug messages
	LogDebug(format string, v ...interface{})

	// LogInfo writes info messages
	LogInfo(format string, v ...interface{})

	// LogError writes error messages
	LogError(format string, v ...interface{})
}

// Creates a logger that writes to the standard output and error streams
func NewStandardLogger() Logger {
	return SimpleLogger{
		InfoLogger:  log.New(os.Stdout, "[INFO]", log.LstdFlags),
		DebugLogger: log.New(os.Stdout, "[DEBUG]", log.LstdFlags),
		ErrorLogger: log.New(os.Stderr, "[ERROR]", log.LstdFlags),
	}
}

// Creates a logger that does nothing with messages
func NewNullLogger() Logger {
	return NullLogger{}
}

// A basic implementation of Logger that writes to stdout and stderr
type SimpleLogger struct {
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	ErrorLogger *log.Logger
}

// LogDebug writes Debug messages to stdout
func (l SimpleLogger) LogDebug(format string, v ...interface{}) {
	l.DebugLogger.Printf(format, v...)
}

// LogInfo writes Info messages to stdout
func (l SimpleLogger) LogInfo(format string, v ...interface{}) {
	l.InfoLogger.Printf(format, v...)
}

// LogError writes to error messages to stderr
func (l SimpleLogger) LogError(format string, v ...interface{}) {
	l.ErrorLogger.Printf(format, v...)
}

// A logger that will do nothing with messages
type NullLogger struct{}

// LogDebug does nothing
func (l NullLogger) LogDebug(format string, v ...interface{}) {
}

// LogInfo does nothing
func (l NullLogger) LogInfo(format string, v ...interface{}) {
}

// LogError does nothing
func (l NullLogger) LogError(format string, v ...interface{}) {
}
