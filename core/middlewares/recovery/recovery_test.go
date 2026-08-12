package recovery

import (
	"context"
	"sync/atomic"
	"testing"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-faster/errors"
	"github.com/gotd/td/bin"
)

type countingInvoker struct {
	calls int32
}

func (i *countingInvoker) Invoke(ctx context.Context, _ bin.Encoder, _ bin.Decoder) error {
	atomic.AddInt32(&i.calls, 1)
	return ctx.Err()
}

func TestHandleStopsRetryWhenInvokeContextIsDone(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	next := &countingInvoker{}
	handler := New(context.Background(), backoff.WithMaxRetries(backoff.NewConstantBackOff(0), 3)).Handle(next)

	err := handler(ctx, nil, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Handle() error = %v, want context.Canceled", err)
	}

	if calls := atomic.LoadInt32(&next.calls); calls != 1 {
		t.Fatalf("Invoke called %d times after context cancellation, want 1", calls)
	}
}
