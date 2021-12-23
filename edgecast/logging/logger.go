// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package logging

import (
	"log"
	"os"
)

const (
	logFlag int = log.Ldate | log.Ltime
)

// Logger defines the the basic implementation that meets the SDK's logging
// needs
type Logger interface {
	// Debug writes debug messages
	Debug(format string, v ...interface{})

	// Info writes info messages
	Info(format string, v ...interface{})

	// Error writes warning messages
	Warn(format string, v ...interface{})

	// Error writes error messages
	Error(format string, v ...interface{})
}

// Creates a logger that writes to a single log file
func NewFileLogger(filePath string) Logger {
	file, err := os.OpenFile(
		filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return SimpleLogger{
		InfoLogger:    log.New(file, "[INFO] ", logFlag),
		DebugLogger:   log.New(file, "[DEBUG] ", logFlag),
		WarningLogger: log.New(file, "[WARN] ", logFlag),
		ErrorLogger:   log.New(file, "[ERROR] ", logFlag),
	}
}

// Creates a logger that writes to the standard output and error streams
func NewStandardLogger() Logger {
	return SimpleLogger{
		InfoLogger:    log.New(os.Stdout, "[INFO] ", logFlag),
		DebugLogger:   log.New(os.Stdout, "[DEBUG] ", logFlag),
		WarningLogger: log.New(os.Stderr, "[WARN] ", logFlag),
		ErrorLogger:   log.New(os.Stderr, "[ERROR] ", logFlag),
	}
}

// Creates a logger that does nothing with messages
func NewNullLogger() Logger {
	return NullLogger{}
}

// A basic implementation of Logger using log.Logger internally
type SimpleLogger struct {
	InfoLogger    *log.Logger
	DebugLogger   *log.Logger
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
}

// Debug writes Debug messages to stdout
func (l SimpleLogger) Debug(format string, v ...interface{}) {
	l.DebugLogger.Printf(format, v...)
}

// Info writes Info messages to stdout
func (l SimpleLogger) Info(format string, v ...interface{}) {
	l.InfoLogger.Printf(format, v...)
}

// Error writes to error messages to stderr
func (l SimpleLogger) Error(format string, v ...interface{}) {
	l.ErrorLogger.Printf(format, v...)
}

// Warn writes to error messages to stderr
func (l SimpleLogger) Warn(format string, v ...interface{}) {
	l.WarningLogger.Printf(format, v...)
}

// A logger that will do nothing with messages
type NullLogger struct{}

// Debug does nothing
func (l NullLogger) Debug(format string, v ...interface{}) {
}

// Info does nothing
func (l NullLogger) Info(format string, v ...interface{}) {
}

// Error does nothing
func (l NullLogger) Error(format string, v ...interface{}) {
}

// Warn does nothing
func (l NullLogger) Warn(format string, v ...interface{}) {
}
