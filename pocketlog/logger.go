package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the requried threshold.
// The defualt output is Stdout
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

func (l *Logger) Debugf(format string, args ...any) {
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold <= LevelDebug {
		l.logf("[Debug] "+format, args...)
	}

}

func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold <= LevelInfo {
		l.logf("[Info] "+format, args...)
	}

}

func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold <= LevelError {
		l.logf("[Error] "+format, args...)
	}
}

func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
