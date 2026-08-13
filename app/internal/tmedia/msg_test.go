package tmedia

import (
	"testing"

	"github.com/gotd/td/tg"
)

func TestGetPhotoInfoSupportsBasicPhotoSize(t *testing.T) {
	item, ok := GetPhotoInfo(&tg.MessageMediaPhoto{
		Photo: &tg.Photo{
			ID: 1001,
			Sizes: []tg.PhotoSizeClass{
				&tg.PhotoSize{Type: "m", Size: 1234},
			},
			DCID: 2,
		},
	})
	if !ok {
		t.Fatal("expected PhotoSize photo info to be supported")
	}

	location, ok := item.InputFileLoc.(*tg.InputPhotoFileLocation)
	if !ok {
		t.Fatalf("input file location = %T, want *tg.InputPhotoFileLocation", item.InputFileLoc)
	}
	if location.ThumbSize != "m" {
		t.Fatalf("thumb type = %q, want %q", location.ThumbSize, "m")
	}
	if item.Size != 1234 {
		t.Fatalf("size = %d, want %d", item.Size, 1234)
	}
}
