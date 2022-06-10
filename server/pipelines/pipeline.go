package pipelines

import (
	"app/server/bots"
	"app/server/database"
	"app/server/models"
	"strconv"
)

func TemplateResponse(b bots.Bot, c models.Client, Message *models.Message) error {
	user := getUserFromMessage(c, *Message)
	state := b.GetState(c.Wid, user)
	if state == "end" {
		b.SetState(b.GetFirstTemplate(), c.Wid, user)
		return nil
	}
	TemplateMessage := b.TemplateMessage(state)
	err := b.SendMessage(TemplateMessage, c.Wid, Message.WidSender)
	return err
}

func ChangeStateBasedOnSelectedOption(b bots.Bot, c models.Client, Message *models.Message) error {
	user := getUserFromMessage(c, *Message)
	Option, err := strconv.Atoi(Message.Message)
	if err != nil && (b.GetState(c.Wid, user) != b.GetFirstTemplate()) {
		b.SendMessage(c.FallbackMessage, c.Wid, Message.WidSender)
		return err
	}
	options := b.GetOptions(c.Wid, user)
	if checkStateOptions(b, c, user, options, Option) {
		if Option != 0 {
			b.SendMessage(c.FallbackMessage, c.Wid, Message.WidSender)
		}
		return nil
	} else {
		if strconv.FormatUint(uint64(c.RateTemplateID), 10) == b.GetState(c.Wid, user) {
			b.RateSession(Option, c.Wid, user)
			b.SendMessage(c.EndMessage, c.Wid, Message.WidSender)
			b.SetState("end", c.Wid, user)
			return err
		}
	}
	b.SetState(b.GetLink(Option, c.Wid, user), c.Wid, user)
	return err
}

func getUserFromMessage(c models.Client, m models.Message) string {
	if c.Wid == m.WidSender {
		return m.WidReceiver
	}
	return m.WidSender
}

func checkStateOptions(b bots.Bot, c models.Client, user string, options []int, Option int) bool {
	if len(options) == 0 {
		return (strconv.FormatUint(uint64(c.RateTemplateID), 10) == b.GetState(c.Wid, user) && (Option < 1 || Option > 3))
	}
	return Option > len(options) || Option < 1
}

func ResetState(b bots.Bot, c models.Client, Message *models.Message) error {
	user := getUserFromMessage(c, *Message)
	b.GetState(c.Wid, user)
	if Message.Message == "reset" {
		b.SetState(b.GetFirstTemplate(), c.Wid, user)
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
