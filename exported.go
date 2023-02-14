package log

import (
	"go.uber.org/zap"
)

type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

func With(args ...interface{}) *zap.SugaredLogger {
	return export.With(args...)
}

func WithError(err error) *zap.SugaredLogger {
	return export.With("err", err)
}

func Log(level Level, args ...interface{}) {
	switch level {
	case InfoLevel:
		export.Info(args...)
	case WarnLevel:
		export.Warn(args...)
	case ErrorLevel:
		export.Error(args...)
	case DPanicLevel:
		export.DPanic(args...)
	case PanicLevel:
		export.Panic(args...)
	case FatalLevel:
		export.Fatal(args...)
	default:
		export.Debug(args...)
	}
}

func Debug(args ...interface{}) {
	export.Debug(args...)
}

func Info(args ...interface{}) {
	export.Info(args...)
}

func Warn(args ...interface{}) {
	export.Warn(args...)
}

func Error(args ...interface{}) {
	export.Error(args...)
}

func DPanic(args ...interface{}) {
	export.DPanic(args...)
}

func Panic(args ...interface{}) {
	export.Panic(args...)
}

func Fatal(args ...interface{}) {
	export.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	export.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	export.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	export.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	export.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	export.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	export.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	export.Fatalf(template, args...)
}
