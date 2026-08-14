package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestCompleteExtFilesFiltersDirectories(t *testing.T) {
	dir := t.TempDir()
	wantFile := filepath.Join(dir, "session.txt")
	if err := os.WriteFile(wantFile, []byte("tdl"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dir, "profiles"), 0o755); err != nil {
		t.Fatal(err)
	}

	files, directive := completeExtFiles("txt")(nil, nil, dir+string(os.PathSeparator))
	if directive != cobra.ShellCompDirectiveFilterDirs {
		t.Fatalf("completion directive = %v; want ShellCompDirectiveFilterDirs", directive)
	}

	if len(files) != 1 || files[0] != wantFile {
		t.Fatalf("completion files = %v; want [%s]", files, wantFile)
	}
}
