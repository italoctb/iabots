package pipelines

import (
	"app/server/bots"
	"app/server/database"
	"app/server/models"
	"strconv"
)

func TemplateResponse(b bots.Bot, Message *models.Message) error {

	state := b.GetState()
	TemplateMessage := b.TemplateMessage(state)
	err := b.SendMessage(TemplateMessage, Message.WidSender)
	return err
}

func ChangeStateBasedOnSelectedOption(b bots.Bot, Message *models.Message) error {
	Option, err := strconv.Atoi(Message.Message)
	if err != nil {
		return err
	}
	options := b.GetOptions()

	if Option > len(options) || Option < 1 {
		b.SendMessage(b.FallbackMessage(), Message.WidReceiver)
		return nil
	}

	b.SetState(b.GetLink(Option))
	return err
}

func ResetState(b bots.Bot, Message *models.Message) error {

	if Message.Message == "reset" {
		b.SetState(b.GetFirstTemplate())
	}
	return nil
}

func ChainProcess(b bots.Bot, Message *models.Message) error {
	ResetState(b, Message)
	currentTemplate := b.GetState()
	ChangeStateBasedOnSelectedOption(b, Message)
	TemplateResponse(b, Message)
	Message.ProcessedAt = true
	Message.FromTemplate = currentTemplate
	db := database.GetDatabase()
	db.Save(&Message)
	return nil
}
