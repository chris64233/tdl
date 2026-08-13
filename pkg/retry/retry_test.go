package retry

import (
	"testing"

	"github.com/gotd/td/tgerr"
)

func TestNewIncludesWorkerBusyTooLongRetry(t *testing.T) {
	mw := New(3)
	r, ok := mw.(retry)
	if !ok {
		t.Fatalf("New returned %T, want retry", mw)
	}

	err := tgerr.New(500, "WORKER_BUSY_TOO_LONG_RETRY")
	if !tgerr.Is(err, r.errors...) {
		t.Fatalf("WORKER_BUSY_TOO_LONG_RETRY should be treated as retryable")
	}
}
