package log

import (
	"context"

	"go.uber.org/zap"
)

type Key int

const contextKey Key = -10

func WithFields(ctx context.Context, fields ...Field) (context.Context, *zap.Logger) {
	cloned := std.zaplog.With(fields...)
	return context.WithValue(ctx, contextKey, cloned), cloned
}

func WithCtx(ctx context.Context) *zap.Logger {
	if ctxLog, ok := ctx.Value(contextKey).(*zap.Logger); ok {
		return ctxLog
	}
	return std.zaplog
}
