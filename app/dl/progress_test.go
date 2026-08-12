package dl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDonePostKeepsTemplateSubdirectory(t *testing.T) {
	dir := t.TempDir()
	nestedDir := filepath.Join(dir, "topic")
	if err := os.MkdirAll(nestedDir, 0o755); err != nil {
		t.Fatal(err)
	}

	tmpFile, err := os.Create(filepath.Join(nestedDir, "file.txt"+tempExt))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpFile.WriteString("downloaded"); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	p := &progress{opts: Options{Dir: dir}}
	elem := &iterElem{to: tmpFile}

	if err := p.donePost(elem); err != nil {
		t.Fatalf("donePost() error = %v", err)
	}

	want := filepath.Join(nestedDir, "file.txt")
	if _, err := os.Stat(want); err != nil {
		t.Fatalf("expected renamed file under template subdirectory: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "file.txt")); !os.IsNotExist(err) {
		t.Fatalf("file was unexpectedly renamed into root download dir, stat err = %v", err)
	}
}
