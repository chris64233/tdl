package recovery

import (
	"context"
	"errors"
	"testing"

	"github.com/cenkalti/backoff/v4"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram"
)

func TestRecoveryStopsWhenRequestContextIsCanceled(t *testing.T) {
	requestCtx, cancel := context.WithCancel(context.Background())
	cancel()

	calls := 0
	next := telegram.InvokeFunc(func(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
		calls++
		return errors.New("transient network failure")
	})

	err := New(context.Background(), backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 1)).
		Handle(next).
		Invoke(requestCtx, nil, nil)
	if err == nil {
		t.Fatal("expected request error")
	}

	if calls != 1 {
		t.Fatalf("expected canceled request context to stop recovery retries after one call, got %d calls", calls)
	}
}
