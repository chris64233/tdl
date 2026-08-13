package extension

import (
	"testing"

	"github.com/fatih/color"
)

func TestNormalizeExtNameRemovesExecutableSuffix(t *testing.T) {
	oldNoColor := color.NoColor
	color.NoColor = true
	t.Cleanup(func() { color.NoColor = oldNoColor })

	got := normalizeExtName("owner/tdl-demo.exe")
	const want = "tdl-demo"
	if got != want {
		t.Fatalf("normalizeExtName() = %q, want %q", got, want)
	}
}
