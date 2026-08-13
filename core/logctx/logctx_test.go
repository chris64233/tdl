package logctx

import (
	"context"
	"testing"
)

func TestFromFallsBackWhenLoggerMissing(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("From should not panic when logger is missing from context: %v", r)
		}
	}()

	if logger := From(context.Background()); logger == nil {
		t.Fatal("expected fallback logger, got nil")
	}
}
