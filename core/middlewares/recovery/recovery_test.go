package recovery

import (
	"context"
	"errors"
	"testing"

	"github.com/cenkalti/backoff/v4"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram"
	"go.uber.org/zap"

	"github.com/iyear/tdl/core/logctx"
)

func TestHandleUsesMiddlewareContextLogger(t *testing.T) {
	rootCtx := logctx.With(context.Background(), zap.NewNop())
	middleware := New(rootCtx, backoff.WithMaxRetries(backoff.NewConstantBackOff(0), 0))

	called := false
	handler := middleware.Handle(telegram.InvokeFunc(func(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
		called = true
		return errors.New("temporary failure")
	}))

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Handle panicked when request context had no logger: %v", r)
		}
		if !called {
			t.Fatal("next invoker was not called")
		}
	}()

	_ = handler.Invoke(context.Background(), nil, nil)
}
