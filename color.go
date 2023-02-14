package log

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

var (
	_levelToColor = map[zapcore.Level]Color{
		zapcore.DebugLevel:  Magenta,
		zapcore.InfoLevel:   Green,
		zapcore.WarnLevel:   Yellow,
		zapcore.ErrorLevel:  Red,
		zapcore.DPanicLevel: Red,
		zapcore.PanicLevel:  Red,
		zapcore.FatalLevel:  Red,
	}
)

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type Color uint8

func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}
