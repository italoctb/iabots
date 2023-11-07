package helper

import (
	"strings"

	"go.mau.fi/whatsmeow/binary/proto"
	protoBuf "google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow"
)

func GetFileOrImageMessage(payload map[string]string, size int, uploaded whatsmeow.UploadResponse) *proto.Message {
	if strings.Contains(payload["mimetype"], "image") {
		return &proto.Message{ImageMessage: &proto.ImageMessage{
			Caption:       protoBuf.String(payload["message"]),
			Url:           protoBuf.String(uploaded.URL),
			DirectPath:    protoBuf.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Mimetype:      protoBuf.String(payload["mimetype"]),
			FileEncSha256: uploaded.FileEncSHA256,
			FileSha256:    uploaded.FileSHA256,
			FileLength:    protoBuf.Uint64(uint64(size)),
		}}
	}
	return &proto.Message{DocumentMessage: &proto.DocumentMessage{
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
}
