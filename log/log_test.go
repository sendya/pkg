package log_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/sendya/pkg/log"
)

func TestCtx(t *testing.T) {
	log.Info("info test..")

	ctx, _ := log.WithFields(context.Background(), log.String("traceId", uuid.New().String()))
	log.WithCtx(ctx).Info("test context", log.Duration("time", time.Second))

	log.Info("test not context", log.Duration("time", time.Microsecond))
}

func TestNop(t *testing.T) {
	c := log.Default()

	log.ResetDefault(log.NewNop())
	log.Info("hidden")

	log.ResetDefault(c)
	log.Info("show log")
}
