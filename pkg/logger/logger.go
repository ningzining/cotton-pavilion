package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"sync"
	"time"
)

type Level = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

type Logger struct {
	mu sync.Mutex
	l  *zap.Logger
	al *zap.AtomicLevel
}

func New(out io.Writer, level Level) *Logger {
	al := zap.NewAtomicLevelAt(level)
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.Format(time.DateTime))
	}

	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config.EncoderConfig),
			//zapcore.NewJSONEncoder(config.EncoderConfig),
			zapcore.AddSync(out),
			al,
		),
		zap.AddCaller(),
		zap.AddStacktrace(ErrorLevel),
		zap.AddCallerSkip(2),
	)

	return &Logger{l: logger, al: &al}
}

func (l *Logger) SetLevel(level Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

var std = New(os.Stdout, InfoLevel)

func Default() *Logger {
	return std
}

func ReplaceDefault(l *Logger) {
	std = l
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

// std logger
func Debug(msg string, fields ...Field) {
	std.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	std.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	std.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	std.Error(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	std.Panic(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	std.Fatal(msg, fields...)
}

func Sync() error {
	return std.Sync()
}
