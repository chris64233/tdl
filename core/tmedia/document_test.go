package tmedia

import (
	"strings"
	"testing"

	"github.com/gotd/td/tg"
)

func getDocumentNameWithoutPanic(t *testing.T, doc *tg.Document) (name string, panicked bool) {
	t.Helper()

	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()

	name = GetDocumentName(doc)
	return name, false
}

func TestGetDocumentNameUsesFilenameAttribute(t *testing.T) {
	name, panicked := getDocumentNameWithoutPanic(t, &tg.Document{
		ID:       42,
		MimeType: "application/octet-stream",
		Attributes: []tg.DocumentAttributeClass{
			&tg.DocumentAttributeFilename{FileName: "report.bin"},
		},
	})
	if panicked {
		t.Fatal("GetDocumentName panicked when filename attribute was present")
	}
	if name != "report.bin" {
		t.Fatalf("GetDocumentName() = %q, want %q", name, "report.bin")
	}
}

func TestGetDocumentNameUsesKnownMimeExtension(t *testing.T) {
	name, panicked := getDocumentNameWithoutPanic(t, &tg.Document{
		ID:       99,
		MimeType: "image/jpeg",
	})
	if panicked {
		t.Fatal("GetDocumentName panicked for known MIME type")
	}
	if !strings.HasPrefix(name, "99.") {
		t.Fatalf("GetDocumentName() = %q, want stable ID prefix", name)
	}
	if name == "99.unknown" {
		t.Fatalf("GetDocumentName() = %q, want known MIME extension", name)
	}
}

func TestGetDocumentNameHandlesEmptyMimeType(t *testing.T) {
	name, panicked := getDocumentNameWithoutPanic(t, &tg.Document{ID: 123})
	if panicked {
		t.Fatal("GetDocumentName panicked for empty MIME type")
	}
	if name != "123.unknown" {
		t.Fatalf("GetDocumentName() = %q, want %q", name, "123.unknown")
	}
}

func TestGetDocumentNameHandlesUnknownMimeType(t *testing.T) {
	name, panicked := getDocumentNameWithoutPanic(t, &tg.Document{ID: 456, MimeType: "application/x-tdl-unknown"})
	if panicked {
		t.Fatal("GetDocumentName panicked for unknown MIME type")
	}
	if name != "456.unknown" {
		t.Fatalf("GetDocumentName() = %q, want %q", name, "456.unknown")
	}
}
