package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	Sugar  *zap.SugaredLogger
	Base   *zap.Logger
	export *zap.SugaredLogger
)

func init() {
	options := getOptions()
	Base = newLog(options)
	Sugar = Base.Sugar()
	export = Sugar.WithOptions(zap.AddCallerSkip(1))
	defer export.Sync()
	Info("日志系统初始化成功")
}

func newLog(options *Options) *zap.Logger {
	production, _ := zap.NewProduction()
	production.Info("asdf")
	cores := make([]zapcore.Core, 0, 2)
	if len(options.File) > 0 {
		logfile, err := os.Create(options.File)
		if err != nil {
			fmt.Printf("日志文件创建失败，%v, %s\n", err, options.File)
		} else {
			encoderConfig := zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				CallerKey:      "call",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}
			cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(logfile), options.GetLevel()))
		}
	}
	if !options.HideConsole {
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			CallerKey:      "call",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		}
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), options.GetLevel()))
	}
	tee := zapcore.NewTee(cores...)
	return zap.New(tee, zap.WithCaller(!options.HideCaller))
}
