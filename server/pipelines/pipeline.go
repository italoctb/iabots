package pipelines

import (
	"app/server/adapters"
	"app/server/database"
	"app/server/models"
	"fmt"
	"strconv"
)

func TemplateResponse(Message *models.Message) error {
	var Session models.Session
	db := database.GetDatabase()
	db.Last(&Session)
	if Session.State == "" {
		return nil
	}
	Template, err := Session.GetActualTemplate(db)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	newMessage := models.Message{Message: Template.GetMessage(),
		WidSender:   "5511989070670",
		WidReceiver: Message.WidSender,
		ProcessedAt: true}
	db.Create(&newMessage)
	Adapter := adapters.Positus{}
	Adapter.SendMessage(newMessage.WidReceiver, newMessage.Message)
	return err
}

func ChangeStateBasedOnSelectedOption(Message *models.Message) error {
	Option, err := strconv.Atoi(Message.Message)
	if err != nil {
		return err
	}
	var Session models.Session
	db := database.GetDatabase()
	db.Last(&Session)
	Template, err := Session.GetActualTemplate(db)
	if Option > len(Template.Options) || Option < 0 {
		newMessage := models.Message{
			Message:     "Opção não existe",
			ProcessedAt: true}
		db.Create(&newMessage)
		return nil
	}
	if Template.Options[Option].Goto == "" {
		newMessage := models.Message{
			Message:     "Opção não está configurada",
			ProcessedAt: true}
		db.Create(&newMessage)
		return nil
	}
	newSession := models.Session{
		State: Template.Options[Option-1].Goto,
	}
	db.Create(&newSession)
	return err
}

func ResetState(Message *models.Message) error {
	db := database.GetDatabase()
	if Message.Message == "reset" {
		newSession := models.Session{
			State: "1"}
		db.Create(&newSession)
	}
	return nil
}

func ChainProcess(Message *models.Message) error {
	ResetState(Message)
	ChangeStateBasedOnSelectedOption(Message)
	TemplateResponse(Message)
	Message.ProcessedAt = true
	db := database.GetDatabase()
	db.Save(&Message)
	return nil
}
