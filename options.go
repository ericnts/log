package log

import (
	"fmt"
	"github.com/ericnts/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	HideCaller  bool   `yaml:"hideCaller"`  // 是否隐藏调用者（打印位置）
	HideConsole bool   `yaml:"hideConsole"` // 是否不输出到控制台
	LazyWrite   bool   `yaml:"lazyWrite"`   // 是否启用异步输出
	SplitTime   uint8  `json:"splitTime"`   // 日志的分割时长，小时
	SplitSize   int    `yaml:"splitSize"`   // 日志分割大小 LogSplit为0时启用
	Level       string `json:"level"`       // 日志级别
	File        string `yaml:"file"`        // 日志保存路径
	PanicFile   string `json:"panicFile"`   // Panic的日志路径
	_level      *zap.AtomicLevel
}

func getOptions() *Options {
	options, err := config.Load[Options]("log")
	if err != nil {
		panic(err)
	}
	return &options
}

func (o *Options) GetLevel() zap.AtomicLevel {
	if o._level == nil {
		level, err := zap.ParseAtomicLevel(o.Level)
		if err != nil {
			fmt.Println("日志等级设定无效")
			level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		}
		o._level = &level
	}
	return *o._level
}
