package adapters

import (
	"strings"
)

type Adapter interface {
	GetUrl() string
	GetToken() string
	SendMessage(widReceiver string, message string) *strings.Reader
}
