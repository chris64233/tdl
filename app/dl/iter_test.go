package dl

import (
	"context"
	"testing"

	"github.com/gotd/td/tg"

	"github.com/iyear/tdl/pkg/tmessage"
)

func TestIterNextHandlesDialogWithoutMessages(t *testing.T) {
	it, err := newIter(nil, nil, [][]*tmessage.Dialog{
		{
			{
				Peer:     &tg.InputPeerUser{UserID: 1},
				Messages: nil,
			},
		},
	}, Options{
		Dir:      t.TempDir(),
		Template: "{{.FileName}}",
	})
	if err != nil {
		t.Fatalf("newIter returned error: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Next panicked for dialog without messages: %v", r)
		}
	}()

	if it.Next(context.Background()) {
		t.Fatal("Next returned true for dialog without messages")
	}
	if err := it.Err(); err != nil {
		t.Fatalf("Next returned unexpected error: %v", err)
	}
}
