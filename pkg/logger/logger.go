package logger

import (
	"context"
	env2 "github.com/rishusahu23/fam-youtube/pkg/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

const (
	UnknownId   = "UNKNOWN_ID"
	CtxTraceKey = "traceIdKey"
	traceIdKey  = "trace_id"
	traceIdTag  = "trace_id"
)

var (
	Log  *zap.Logger
	once sync.Once

	IsDevLogger bool

	atomicLevel zap.AtomicLevel
)

func Init(env string) {
	once.Do(func() {
		initLogger(env)
	})
}

func initLogger(env string) {
	var err error
	switch env {
	case env2.StagingEnv, env2.UatEnv, env2.TestEnv:
		Log, atomicLevel, err = NewProductionWithDebug()
	default:
		Log, atomicLevel, err = NewProduction()
	}

	if err != nil {
		log.Panicf("Failed to initiate zap logger: %v", err)
	}
}

func NewProductionWithDebug(options ...zap.Option) (*zap.Logger, zap.AtomicLevel, error) {
	config := zap.NewProductionConfig()
	config.Development = true
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	options = append(options, zap.AddCallerSkip(1))

	logger, err := config.Build(options...)
	if err != nil {
		return nil, zap.AtomicLevel{}, err
	}

	return logger, config.Level, nil
}

func NewProduction(options ...zap.Option) (*zap.Logger, zap.AtomicLevel, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	options = append(options, zap.AddCallerSkip(1))

	logger, err := config.Build(options...)
	if err != nil {
		return nil, zap.AtomicLevel{}, err
	}

	return logger, config.Level, nil
}

func GetCtxLogFields(ctx context.Context, fields ...zap.Field) []zap.Field {
	contextKeys := []string{traceIdKey}

	for _, key := range contextKeys {
		switch key {
		case traceIdTag:
			fields = WithTraceId(ctx, fields...)
		}
	}
	return fields
}

func WithTraceId(ctx context.Context, fields ...zap.Field) []zap.Field {
	return appendFieldIfAbsent(fields, zap.String("traceIdTag", TraceIdFromContext(ctx)))
}

func appendFieldIfAbsent(fields []zap.Field, elem zap.Field) []zap.Field {
	for _, field := range fields {
		if field.Key == elem.Key {
			return fields
		}
	}

	return append(fields, elem)
}

func TraceIdFromContext(ctx context.Context) string {
	id := fromContextOrGRPCMetadata(ctx, traceIdKey, CtxTraceKey)
	if id != nil {
		return id.(string)
	}
	return UnknownId
}

func fromContextOrGRPCMetadata(ctx context.Context, metadataKey string, ctxKey interface{}) interface{} {
	return fromContext(ctx, ctxKey)
}

func fromContext(ctx context.Context, key interface{}) interface{} {
	if ctx.Value(key) != nil {
		return ctx.Value(key)
	}
	return nil
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	Log.Info(msg, GetCtxLogFields(ctx, fields...)...)

}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	Log.Debug(msg, GetCtxLogFields(ctx, fields...)...)
}
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	Log.WithOptions(zap.AddStacktrace(zap.ErrorLevel)).Error(msg, GetCtxLogFields(ctx, fields...)...)
}

func WarnWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	Log.Warn(msg, GetCtxLogFields(ctx, fields...)...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

func Dpanic(msg string, fields ...zap.Field) {
	Log.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Log.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

func PanicWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	Log.Panic(msg, GetCtxLogFields(ctx, fields...)...)
}

func FatalWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	Log.Fatal(msg, GetCtxLogFields(ctx, fields...)...)
}

func DebugNoCtx(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func InfoNoCtx(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func ErrorNoCtx(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func RecoverPanicAndError(ctx context.Context) {
	if r := recover(); r != nil {
		Log.WithOptions(zap.AddStacktrace(zap.ErrorLevel)).Error("recovered from panic", GetCtxLogFields(ctx, zap.Any("recoverMsg", r))...)
	}
}
