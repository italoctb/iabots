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
	if strconv.FormatUint(uint64(c.RateTemplateID), 10) == state {
		b.SetState(b.GetFirstTemplate(), c.Wid)
		return nil
	}
	err := b.SendMessage(TemplateMessage, c.Wid, Message.WidSender)
	return err
}

func ChangeStateBasedOnSelectedOption(b bots.Bot, c models.Client, Message *models.Message) error {
	Option, err := strconv.Atoi(Message.Message)
	if err != nil {
		return err
	}
	options := b.GetOptions()

	if checkStateOptions(b, c, options, Option) {
		b.SendMessage(c.FallbackMessage, c.Wid, Message.WidReceiver)
		return nil
	} else {
		if strconv.FormatUint(uint64(c.RateTemplateID), 10) == b.GetState() {
			b.RateSession(Option)
			b.SendMessage(c.EndMessage, c.Wid, Message.WidReceiver)
			return err
		}
	}
	b.SetState(b.GetLink(Option), c.Wid)
	return err
}

func checkStateOptions(b bots.Bot, c models.Client, options []int, Option int) bool {
	if len(options) == 0 {
		return (strconv.FormatUint(uint64(c.RateTemplateID), 10) == b.GetState() && (Option < 1 || Option > 3))
	}
	return Option > len(options) || Option < 1
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
