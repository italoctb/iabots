package bots

type Bot interface {
	FallbackMessage() string
	GetFirstTemplate() string
	EndMessage() string
	SendMessage(message string) error
	GetOptions() []int
	GetLink(int) string
	TemplateMessage(state string) string
	SetState(string) string
	GetState() string
}
