package retry

import (
	"context"
	"testing"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tgerr"
	"go.uber.org/zap"

	"github.com/iyear/tdl/pkg/logger"
)

func TestRetryTreatsMemoryLimitExitAsInternalError(t *testing.T) {
	ctx := logger.With(context.Background(), zap.NewNop())
	attempts := 0

	handler := New(2).Handle(telegram.InvokeFunc(func(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
		attempts++
		if attempts == 1 {
			return tgerr.New(500, "memory limit exit")
		}
		return nil
	}))

	if err := handler.Invoke(ctx, nil, nil); err != nil {
		t.Fatalf("Invoke() error = %v, want nil after retry", err)
	}
	if attempts != 2 {
		t.Fatalf("attempts = %d, want 2", attempts)
	}
}
