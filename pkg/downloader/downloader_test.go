package downloader

import (
	"context"
	"errors"
	"testing"
)

func TestDownloadReturnsCanceledContextBeforeFileWork(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	d := New(nil, 1, 1, nil)
	err := d.download(ctx, &Item{Name: "bad\x00name.bin"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("download() error = %v; want context.Canceled", err)
	}
}
