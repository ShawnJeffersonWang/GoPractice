package log

import (
	"fmt"
	"io"
	"os"
)

// Verbose decides whether the informatc logs should be output.
var Verbose bool

// Logger .
type Logger struct {
	Println func(w io.Writer, a ...interface{}) (n int, err error)
	Printf  func(w io.Writer, format string, a ...interface{}) (n int, err error)
}

var defaultLogger = Logger{
	Println: fmt.Fprintln,
	Printf:  fmt.Fprintf,
}

// SetDefaultLogger sets the default logger.
func SetDefaultLogger(l Logger) {
	defaultLogger = l
}

// Warn .
func Warn(v ...interface{}) {
	defaultLogger.Println(os.Stderr, v...)
}

// Warnf .
func Warnf(format string, v ...interface{}) {
	defaultLogger.Printf(os.Stderr, format, v...)
}

// Info .
func Info(v ...interface{}) {
	if Verbose {
		defaultLogger.Println(os.Stderr, v...)
	}
}

// Infof .
func Infof(format string, v ...interface{}) {
	if Verbose {
		defaultLogger.Printf(os.Stderr, format, v...)
	}
}
