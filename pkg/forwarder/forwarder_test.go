package forwarder

import (
	"testing"

	"github.com/gotd/td/tg"
)

func TestMediaSizeSumAllowsTextOnlyMessage(t *testing.T) {
	total, err := mediaSizeSum(&tg.Message{
		ID:      101,
		Message: "text only",
	})
	if err != nil {
		t.Fatalf("mediaSizeSum returned error for text-only message: %v", err)
	}
	if total != 0 {
		t.Fatalf("mediaSizeSum total = %d, want 0", total)
	}
}
