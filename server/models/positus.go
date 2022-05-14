package models

import (
	"fmt"
	"strings"
)

type Positus struct {
	Url   string
	Token string
}

func CreatePositusObject(url string, token string) Positus {
	var positusObject Positus
	positusObject.Url = url
	positusObject.Token = token
	return positusObject
}

func (p Positus) TextMessage(widReceiver string, message string, method string) *strings.Reader {
	payload := strings.NewReader("{\n" +
		"	\"to\": \"" + widReceiver + "\",\n" +
		"	\"type\": \"text\",\n" +
		"	\"text\": {\n" +
		"		\"body\": \"" + message + "\"\n" +
		"	}" +
		"}")
	fmt.Println(payload)
	return payload
}
