package tmedia

import (
	"reflect"
	"testing"

	"github.com/gotd/td/tg"
)

func TestMediaInfoRetainsUploadDate(t *testing.T) {
	const uploadDate = 1710000000

	photoMedia, ok := GetPhotoInfo(&tg.MessageMediaPhoto{
		Photo: &tg.Photo{
			ID:   123,
			Date: uploadDate,
			Sizes: []tg.PhotoSizeClass{
				&tg.PhotoSize{Type: "x", Size: 42},
			},
		},
	})
	if !ok {
		t.Fatal("GetPhotoInfo returned false")
	}
	assertMediaDate(t, photoMedia, uploadDate)

	documentMedia, ok := GetDocumentInfo(&tg.MessageMediaDocument{
		Document: &tg.Document{
			ID:       456,
			Date:     uploadDate,
			Size:     64,
			MimeType: "application/pdf",
			Attributes: []tg.DocumentAttributeClass{
				&tg.DocumentAttributeFilename{FileName: "report.pdf"},
			},
		},
	})
	if !ok {
		t.Fatal("GetDocumentInfo returned false")
	}
	assertMediaDate(t, documentMedia, uploadDate)
}

func assertMediaDate(t *testing.T, media *Media, want int) {
	t.Helper()

	field := reflect.ValueOf(media).Elem().FieldByName("Date")
	if !field.IsValid() {
		t.Fatal("media info should retain the Telegram upload date")
	}
	if got := field.Int(); got != int64(want) {
		t.Fatalf("media date = %d, want %d", got, want)
	}
}
