package tmedia

import (
	"testing"

	"github.com/gotd/td/tg"
)

func TestGetPhotoInfoUsesStablePhotoIDName(t *testing.T) {
	first, ok := GetPhotoInfo(&tg.MessageMediaPhoto{
		Photo: &tg.Photo{
			ID: 1001,
			Sizes: []tg.PhotoSizeClass{
				&tg.PhotoSize{Type: "x", Size: 123},
			},
			DCID: 2,
		},
	})
	if !ok {
		t.Fatal("expected first photo info")
	}

	second, ok := GetPhotoInfo(&tg.MessageMediaPhoto{
		Photo: &tg.Photo{
			ID: 1002,
			Sizes: []tg.PhotoSizeClass{
				&tg.PhotoSize{Type: "x", Size: 456},
			},
			DCID: 2,
		},
	})
	if !ok {
		t.Fatal("expected second photo info")
	}

	if first.Name != "1001.jpg" {
		t.Fatalf("first photo name = %q, want %q", first.Name, "1001.jpg")
	}
	if second.Name != "1002.jpg" {
		t.Fatalf("second photo name = %q, want %q", second.Name, "1002.jpg")
	}
	if first.Name == second.Name {
		t.Fatalf("photo names must be unique, both got %q", first.Name)
	}
}
