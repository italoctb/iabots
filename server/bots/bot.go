package bots

import (
	"app/server/adapters"
	"app/server/database"
	"app/server/models"
	"fmt"
	"strconv"
)

type ExampleBot struct {
}

func (l ExampleBot) SendMessage(message string, sender string, receiver string) error {
	db := database.GetDatabase()
	newMessage := models.Message{
		Message:     message,
		WidSender:   sender,
		WidReceiver: receiver,
		ProcessedAt: true}
	db.Create(&newMessage)
	Positus := adapters.Positus{}
	err := Positus.SendMessage(receiver, message)
	return err
}

func (l ExampleBot) SetState(link string, widClient string, widUser string) string {
	db := database.GetDatabase()
	newSession := models.Session{
		State:     link,
		WidClient: widClient,
		WidUser:   widUser,
	}
	db.Create(&newSession)
	return ""
}

func (l ExampleBot) GetState(widClient string, widUser string) string {
	var Session models.Session
	db := database.GetDatabase()
	db.Where("wid_client = ? AND wid_user = ?", widClient, widUser).Last(&Session)
	fmt.Print("Sessão: " + Session.State)
	if Session.State == "" {
		l.SetState(l.GetFirstTemplate(), widClient, widUser)
		return l.GetFirstTemplate()
	}
	return Session.State
}

func (l ExampleBot) GetFirstTemplate() string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").First(&Template)
	return strconv.FormatUint(uint64(Template.ID), 10)
}

func (l ExampleBot) GetOptions(widClient string, widUser string) []int {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", l.GetState(widClient, widUser))

	list := []int{}
	pivot := 0
	for _, o := range Template.Options {
		list = append(list, pivot)
		pivot += 1
		fmt.Println(o)
	}
	return list
}

func (l ExampleBot) GetLink(position int, widClient string, widUser string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", l.GetState(widClient, widUser))
	return Template.Options[position-1].Goto
}

func (l ExampleBot) TemplateMessage(state string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", state)
	return Template.GetMessage()
}

func (l ExampleBot) RateSession(rate int, widClient string, widUser string) {
	var Session models.Session
	db := database.GetDatabase()
	db.Where("wid_client = ? AND wid_user = ?", widClient, widUser).Last(&Session)
	Session.Rate = rate
	db.Save(&Session)
}
