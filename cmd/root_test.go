package cmd

import (
	"testing"

	"github.com/iyear/tdl/pkg/consts"
)

func TestRootCommandPartSizeDefault(t *testing.T) {
	flag := New().PersistentFlags().Lookup(consts.FlagPartSize)
	if flag == nil {
		t.Fatalf("expected %q flag to be registered", consts.FlagPartSize)
	}

	const want = "524288"
	got := flag.DefValue
	if got != want {
		t.Fatalf("unexpected default %q for %q flag, want %q", got, consts.FlagPartSize, want)
	}
}
