package logger

import (
	"context"
	"encoding/json"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LogIDKey = "log_id"
)

var logger *zap.Logger

func Init(logFile string) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic(err)
	}

	if logFile == "" {
		logFile = "logs/logger.log"
	}
	rotateConfig := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // compress rotated files
	}

	encConfig := zap.NewProductionEncoderConfig()
	encConfig.TimeKey = "timestamp"
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(encConfig)
	fileWS := zapcore.AddSync(rotateConfig)
	logLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	logger = zap.New(
		zapcore.NewCore(fileEncoder, fileWS, logLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	defer logger.Sync()
}

func initConsoleLogger() {
	logger = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(zap.PanicLevel),
		),
	)
}

func getLogIDField(ctx context.Context) zap.Field {
	logID := ctx.Value(LogIDKey)
	if logID == nil {
		return zap.String(LogIDKey, "unknown")
	}
	return zap.String(LogIDKey, logID.(string))
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	if logger == nil {
		initConsoleLogger()
	}

	allFields := append([]zap.Field{getLogIDField(ctx)}, fields...)
	logger.Info(msg, allFields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	if logger == nil {
		initConsoleLogger()
	}

	allFields := append([]zap.Field{getLogIDField(ctx)}, fields...)
	logger.Warn(msg, allFields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	if logger == nil {
		initConsoleLogger()
	}

	allFields := append([]zap.Field{getLogIDField(ctx)}, fields...)
	logger.Error(msg, allFields...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if logger == nil {
		initConsoleLogger()
	}

	allFields := append([]zap.Field{getLogIDField(ctx)}, fields...)
	logger.Fatal(msg, allFields...)
}

func ErrorLog(err error) zap.Field {
	return zap.Error(err)
}

func EmplaceKV[T any](key string, val T) zap.Field {
	valByte, err := json.Marshal(val)
	valStr := string(valByte)
	if err != nil {
		valStr = "failed to marshal"
	}

	return zap.String(key, valStr)
}
