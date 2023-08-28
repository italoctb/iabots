package bots

import "app/server/models"

type Bot interface {
	GetFirstTemplate(string) string
	SendMessage(message string, sender string, receiver string) error
	GetOptions(string, string) []int
	GetLink(int, string, string) string
	TemplateMessage(state string) string
	SetState(string, string, string) string
	GetSession(string, string) models.Session
	GetStateTemplate(string, string) string
	RateSession(rate int, client string, user string)
}
