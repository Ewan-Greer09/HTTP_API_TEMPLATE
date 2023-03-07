package logger

import (
	"log"
	"os"
)

// This package will be used to log errors and other
// information to a file, database or service

// TODO: Add text colouring to the output
type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warnLogger.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.errorLogger.Println(v...)
}

func (l *Logger) SetFlags(flag int) {
	l.infoLogger.SetFlags(flag)
	l.warnLogger.SetFlags(flag)
	l.errorLogger.SetFlags(flag)
}

func (l *Logger) SetOutput(w *os.File) {
	l.infoLogger.SetOutput(w)
	l.warnLogger.SetOutput(w)
	l.errorLogger.SetOutput(w)
}
