package bots

type Bot interface {
	GetFirstTemplate(string) string
	SendMessage(message string, sender string, receiver string) error
	GetOptions(string, string) []int
	GetLink(int, string, string) string
	TemplateMessage(state string) string
	SetState(string, string, string) string
	GetState(string, string) string
	RateSession(rate int, client string, user string)
}
