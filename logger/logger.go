package logger

import (
	"log"
	"os"
)

// This package will be used to log errors and other
// information to a file, database or service
type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	output      *os.File
}

// NewLogger creates a new logger
func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
		output:      os.Stdout,
	}
}

// Info logs an info message to the info logger
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Warn logs a warning to the warning logger
func (l *Logger) Warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

// Error logs an error to the error logger
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

// Infof logs an info message to the info logger
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Warnf logs a warning to the warning logger
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warnLogger.Printf(format, v...)
}

// Errorf logs an error to the error logger
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Infoln logs an info message to the info logger
func (l *Logger) Infoln(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Warnln logs a warning to the warning logger
func (l *Logger) Warnln(v ...interface{}) {
	l.warnLogger.Println(v...)
}

// Errorln logs an error to the error logger
func (l *Logger) Errorln(v ...interface{}) {
	l.errorLogger.Println(v...)
}

// SetFlags sets the output flags for the logger
func (l *Logger) SetFlags(flag int) {
	l.infoLogger.SetFlags(flag)
	l.warnLogger.SetFlags(flag)
	l.errorLogger.SetFlags(flag)
}

// SetOutput sets the output destination for the logger
func (l *Logger) SetOutput(w string) {
	file, err := os.OpenFile(w, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		l.errorLogger.Println(err)
	}

	l.infoLogger.SetOutput(file)
	l.warnLogger.SetOutput(file)
	l.errorLogger.SetOutput(file)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.errorLogger.Fatal(v...)
}
