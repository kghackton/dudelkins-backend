package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.DebugLevel
}

func init() {
	level := getLoggerLevel("debug")

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	zapLogger, err := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(zap.WithCaller(true), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	globalLogger = zapLogger.Sugar()
}

func Sync() (err error) {
	return globalLogger.Sync()
}

// DEBUG --

func Debug(args ...interface{}) {
	globalLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	globalLogger.Debugf(template, args...)
}
func Debugw(msg string, keysAndValues ...interface{}) {
	globalLogger.Debugw(msg, keysAndValues...)
}

// INFO --

func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	globalLogger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	globalLogger.Infow(msg, keysAndValues...)
}

// WARN --

func Warn(args ...interface{}) {
	globalLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	globalLogger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	globalLogger.Warnw(msg, keysAndValues...)
}

// ERROR --

func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	globalLogger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	globalLogger.Errorw(msg, keysAndValues...)
}

// DPANIC --

func DPanic(args ...interface{}) {
	globalLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	globalLogger.DPanicf(template, args...)
}

// PANIC --

func Panic(args ...interface{}) {
	globalLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	globalLogger.Panicf(template, args...)
}

// FATAL --

func Fatal(args ...interface{}) {
	globalLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	globalLogger.Fatalf(template, args...)
}
