package logger

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-iot-platform/lib/structs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// DefautlLogBucket default directory to log
	DefautlLogBucket = "log"
	// DefautlLogLocation default location to log
	DefautlLogLocation = "./logs"
	// DefaultServiceName default service name in log
	DefaultServiceName = "default"
)

// A Logger represents an active logging object that generates lines
// of JSON output to an io.Writer. Each logging operation makes a single
// call to the Writer's Write method. There is no guaranty on access
// serialization to the Writer. If your Writer is not thread safe,
// you may consider a sync wrapper.
type Logger struct {
	w       zerolog.LevelWriter
	level   zerolog.Level
	sampler zerolog.Sampler
	context []byte
	hooks   []zerolog.Hook
}

// Context configures a new sub-logger with contextual fields.
type Context struct {
	l Logger
}

// Init config log
func Init(serviceName string, logLevel int, logLocation string, logBucket string) (*Logger, *os.File, error) {
	var f *os.File
	prodLog := true
	if logBucket == "" {
		logBucket = DefautlLogBucket
	}
	if logLocation == "" {
		prodLog = false
		logLocation = DefautlLogLocation
	}
	if serviceName == "" {
		serviceName = DefaultServiceName
	}
	// Init logger
	currentTime := time.Now()
	filePath := fmt.Sprintf("%s/%s-%s.log", logLocation, logBucket, currentTime.Format("02-01-2006"))
	if _, err := os.Stat(logLocation); err != nil {
		if os.IsNotExist(err) {
			// path/to/whatever does *not* exist
			err = os.Mkdir(logLocation, 0755)
			if err != nil {
				return nil, f, err
			}

		} else {
			// Schrodinger: file may or may not exist. See err for details.

			// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
			return nil, nil, errors.New("file error")
		}
	}
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, f, err
	}

	logger, err := config(logLevel, f, serviceName, prodLog)
	if err != nil {
		return nil, f, err
	}
	return logger, f, nil
}

// config - custom time format for logger of empty string to use default
func config(lvl int, file *os.File, serviceName string, prodLog bool) (*Logger, error) {
	var logLevel zerolog.Level
	logger := &Logger{}
	//! File
	if prodLog == true {
		log.Logger = log.Output(file)
	}
	//!
	switch lvl {
	case -1:
		logLevel = zerolog.TraceLevel
	case 0:
		logLevel = zerolog.DebugLevel
	case 1:
		logLevel = zerolog.InfoLevel
	case 2:
		logLevel = zerolog.WarnLevel
	case 3:
		logLevel = zerolog.ErrorLevel
	case 4:
		logLevel = zerolog.FatalLevel
	case 5:
		logLevel = zerolog.PanicLevel
	default:
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	log.Logger = log.With().Caller().Str("service", serviceName).Logger()
	structs.Merge(logger, log.Logger)

	return logger, nil
}

// Output duplicates the global logger and sets w as its output.
func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return log.Output(w)
}

// With creates a child logger with the field added to its context.
func (l *Logger) With() zerolog.Context {
	return log.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func (l *Logger) Level(level zerolog.Level) zerolog.Logger {
	return log.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return log.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return log.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Err(err error) *zerolog.Event {
	return log.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Trace() *zerolog.Event {
	return log.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Debug() *zerolog.Event {
	return log.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Info() *zerolog.Event {
	return l.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Warn() *zerolog.Event {
	return log.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Error() *zerolog.Event {
	return log.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Fatal() *zerolog.Event {
	return log.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Panic() *zerolog.Event {
	return log.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return log.WithLevel(level)
}

// Log starts a new message with no level. Setting GlobalLevel to
// Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Log() *zerolog.Event {
	return log.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	log.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
