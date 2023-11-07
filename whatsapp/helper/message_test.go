package helper

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	protoBuf "google.golang.org/protobuf/proto"
)

func TestGetFileOrImageMessage(t *testing.T) {
	payload := map[string]string{
		"message":  "test",
		"mimetype": "image/jpeg",
	}
	uploaded := whatsmeow.UploadResponse{
		URL:           "https://test.com",
		DirectPath:    "https://test.com",
		MediaKey:      []byte("test"),
		FileEncSHA256: []byte("test"),
		FileSHA256:    []byte("test"),
		FileLength:    1,
	}
	size := 1
	message := GetFileOrImageMessage(payload, size, uploaded)
	correctMsg := &proto.Message{ImageMessage: &proto.ImageMessage{
		Caption:       protoBuf.String(payload["message"]),
		Url:           protoBuf.String(uploaded.URL),
		DirectPath:    protoBuf.String(uploaded.DirectPath),
		MediaKey:      uploaded.MediaKey,
		Mimetype:      protoBuf.String(payload["mimetype"]),
		FileEncSha256: uploaded.FileEncSHA256,
		FileSha256:    uploaded.FileSHA256,
		FileLength:    protoBuf.Uint64(uint64(size)),
	}}
	if message == nil {
		t.Error("Message is nil")
	}
	assert.Equal(t, message, correctMsg)

	payload = map[string]string{
		"message":  "test",
		"mimetype": "application/pdf",
		"title":    "test",
	}
	message = GetFileOrImageMessage(payload, size, uploaded)
	correctMsg = &proto.Message{DocumentMessage: &proto.DocumentMessage{
		Caption:       protoBuf.String(payload["message"]),
		Title:         protoBuf.String(payload["title"]),
		FileName:      protoBuf.String(payload["title"]),
		Url:           protoBuf.String(uploaded.URL),
		DirectPath:    protoBuf.String(uploaded.DirectPath),
		MediaKey:      uploaded.MediaKey,
		Mimetype:      protoBuf.String(payload["mimetype"]),
		FileEncSha256: uploaded.FileEncSHA256,
		FileSha256:    uploaded.FileSHA256,
		FileLength:    protoBuf.Uint64(uint64(size)),
	}}
	if message == nil {
		t.Error("Message is nil")
	}
	assert.Equal(t, message, correctMsg)
}
