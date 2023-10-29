package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
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
	l  *zap.Logger
	al *zap.AtomicLevel
}

func New(infoOut io.Writer, errOut io.Writer, level Level) *Logger {
	if infoOut == nil {
		infoOut = os.Stdout
	}
	if errOut == nil {
		errOut = os.Stderr
	}

	al := zap.NewAtomicLevelAt(level)
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.Format(time.DateTime))
	}

	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= al.Level()
	})
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level > al.Level()
	})

	logger := zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(config.EncoderConfig),
				zapcore.AddSync(infoOut),
				infoLevel,
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(config.EncoderConfig),
				zapcore.AddSync(errOut),
				errorLevel,
			),
		),
		zap.AddCaller(),
		zap.AddStacktrace(errorLevel),
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

var std = New(os.Stdout, os.Stdout, InfoLevel)

func Default() *Logger {
	return std
}

func ReplaceDefault(l *Logger) {
	std = l
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

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
