package config

import (
	"io"
	"log"
	"os"
)

// Create a Logger Struct Type
type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

// Create a function to create a new logger
func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime|log.Lmsgprefix)

	return &Logger{
		debug:   log.New(writer, "\033[1;32mDEBUG >\033[0m ", logger.Flags()),
		info:    log.New(writer, "\033[1;36mINFO >\033[0m ", logger.Flags()),
		warning: log.New(writer, "\033[1;93mWARNING >\033[0m ", logger.Flags()),
		err:     log.New(writer, "\033[1;91mERROR >\033[0m ", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Formatted Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
