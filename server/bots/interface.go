package bots

type Bot interface {
	FallbackMessage(message string) string
	GetFirstTemplate() string
	EndMessage(message string) string
	SendMessage(message string, sender string, receiver string) error
	GetOptions() []int
	GetLink(int) string
	TemplateMessage(state string) string
	SetState(string, string) string
	GetState() string
}
