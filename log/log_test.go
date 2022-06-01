package log_test

import (
	"context"
	"os"
	"runtime"
	"testing"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/sendya/pkg/ginx/traceid"
	"github.com/sendya/pkg/log"
)

func TestCtx(t *testing.T) {
	log.Info("info test..")

	ctx, _ := log.WithFields(context.Background(), log.String("traceId", traceid.NewID()))
	log.WithCtx(ctx).Info("test context", log.Int("goroutine", runtime.NumGoroutine()), log.Duration("time", time.Second))

	// 协程里带上下文日志
	go func() {
		log.WithCtx(ctx).Info("test content222", log.Int("goroutine", runtime.NumGoroutine()))
	}()

	time.Sleep(time.Second)

	log.Info("test not context", log.Duration("time", time.Microsecond), log.Int("goroutine", runtime.NumGoroutine()))
}

func TestJSONEncoder(t *testing.T) {
	encoder := log.NewProductionEncoderConfig()
	encoder.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
		zapcore.CapitalLevelEncoder(l, pae)
	}
	encoder.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("15:04:05.000"))
	}
	log.ResetDefault(log.NewWithEncoder(os.Stdout, log.DebugLevel, encoder, log.WithCaller(true)))

	log.Info("test not context", log.Duration("time", time.Microsecond))
}

func TestCustomLogger(t *testing.T) {
	core := log.NewCore(
		log.NewJSONEncoder(log.NewProductionEncoderConfig()),
		log.AddSync(os.Stdout),
		log.InfoLevel,
	)
	log.ResetDefault(log.NewCustom(core))

	log.Info("test not context", log.Duration("time", time.Microsecond))
}

func TestNop(t *testing.T) {
	c := log.Default()

	log.ResetDefault(log.NewNop())
	log.Info("hidden")

	log.ResetDefault(c)
	log.Info("show log")
}
