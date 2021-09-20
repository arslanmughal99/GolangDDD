package logger

import (
	"fmt"
	"time"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _logger *zap.Logger

func Init() {
	lg := zap.NewDevelopmentEncoderConfig()

	lg.EncodeTime = customTimeEncoder
	lg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	_logger = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(lg),
			zapcore.AddSync(colorable.NewColorableStdout()),
			zapcore.DebugLevel,
		),
	)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[ %s ]", t.Format("01/02/06 3:04 PM")))
}

func Warn(msg string, value interface{}) {
	_logger.Warn(msg, zap.Any("Details", value))
}

func Info(msg string, value interface{}) {
	_logger.Info(msg, zap.Any("Details", value))
}

func Debug(msg string, value interface{}) {
	_logger.Debug(msg, zap.Any("Details", value))
}

func Error(msg string, e error) {
	_logger.Error(msg, zap.Error(e))
}

func Fatal(msg string, value interface{}) {
	_logger.Fatal(msg, zap.Any("Details", value))
}

func Flush() {
	if err := _logger.Sync(); err != nil {
		panic("Failed to flush logs")
	}
}
