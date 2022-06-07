package pipelines

import (
	"app/server/bots"
	"app/server/database"
	"app/server/models"
	"strconv"
)

func TemplateResponse(b bots.Bot, c models.Client, Message *models.Message) error {

	state := b.GetState()
	TemplateMessage := b.TemplateMessage(state)
	err := b.SendMessage(TemplateMessage, c.Wid, Message.WidSender)
	return err
}

func ChangeStateBasedOnSelectedOption(b bots.Bot, c models.Client, Message *models.Message) error {
	Option, err := strconv.Atoi(Message.Message)
	if err != nil {
		return err
	}
	options := b.GetOptions()

	if Option > len(options) || Option < 1 {
		b.SendMessage(b.FallbackMessage(c.FallbackMessage), c.Wid, Message.WidReceiver)
		return nil
	}

	b.SetState(b.GetLink(Option), c.Wid)
	return err
}

func ResetState(b bots.Bot, c models.Client, Message *models.Message) error {

	if Message.Message == "reset" {
		b.SetState(b.GetFirstTemplate(), c.Wid)
	}
	return nil
}

func ChainProcess(b bots.Bot, c models.Client, Message *models.Message) error {
	ResetState(b, c, Message)
	ChangeStateBasedOnSelectedOption(b, c, Message)
	TemplateResponse(b, c, Message)
	Message.ProcessedAt = true
	db := database.GetDatabase()
	db.Save(&Message)
	return nil
}
