package bots

type Bot interface {
	GetFirstTemplate() string
	SendMessage(message string, sender string, receiver string) error
	GetOptions() []int
	GetLink(int) string
	TemplateMessage(state string) string
	SetState(string, string) string
	GetState() string
	RateSession(rate int)
}
