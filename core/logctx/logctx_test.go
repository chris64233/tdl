package logctx

import (
	"context"
	"testing"

	"go.uber.org/zap"
)

func getLoggerWithoutPanic(ctx context.Context) (logger *zap.Logger, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()

	return From(ctx), false
}

func TestFromReturnsStoredLogger(t *testing.T) {
	expected := zap.NewNop().Named("stored")
	got, panicked := getLoggerWithoutPanic(With(context.Background(), expected))
	if panicked {
		t.Fatal("From panicked when logger was stored in context")
	}
	if got != expected {
		t.Fatal("From did not return logger stored in context")
	}
}

func TestFromHandlesContextWithoutLogger(t *testing.T) {
	logger, panicked := getLoggerWithoutPanic(context.Background())
	if panicked {
		t.Fatal("From panicked when context had no logger")
	}
	if logger == nil {
		t.Fatal("From returned nil for context without logger")
	}
}

func TestNamedHandlesContextWithoutLogger(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Fatal("Named panicked when context had no logger")
		}
	}()

	ctx := Named(context.Background(), "fallback")
	if logger := From(ctx); logger == nil {
		t.Fatal("Named stored nil logger")
	}
}
