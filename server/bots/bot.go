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

func (l ExampleBot) SetState(link string, widCostumer string, widUser string) string {
	db := database.GetDatabase()
	newSession := models.Session{
		State:       link,
		WidCostumer: widCostumer,
		WidUser:     widUser,
	}
	db.Create(&newSession)
	return ""
}

func (l ExampleBot) GetState(widCostumer string, widUser string) string {
	var Session models.Session
	db := database.GetDatabase()
	fmt.Print("wid costumer = " + widCostumer)
	db.Where("wid_costumer = ? AND wid_user = ?", widCostumer, widUser).Last(&Session)
	fmt.Print("Sessão: " + Session.State)
	if Session.State == "" || Session.State == "0" {
		l.SetState(l.GetFirstTemplate(widCostumer), widCostumer, widUser)
		return l.GetFirstTemplate(widCostumer)
	}
	return Session.State
}

func (l ExampleBot) GetFirstTemplate(widCostumer string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "wid = ? AND is_first=?", widCostumer, true)
	return strconv.FormatUint(uint64(Template.ID), 10)
}

func (l ExampleBot) GetOptions(widCostumer string, widUser string) []int {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", l.GetState(widCostumer, widUser))

	list := []int{}
	pivot := 0
	for _, o := range Template.Options {
		list = append(list, pivot)
		pivot += 1
		fmt.Println(o)
	}
	return list
}

func (l ExampleBot) GetLink(position int, widCostumer string, widUser string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", l.GetState(widCostumer, widUser))
	return Template.Options[position-1].Goto
}

func (l ExampleBot) TemplateMessage(state string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", state)
	return Template.GetMessage()
}

func (l ExampleBot) RateSession(rate int, widCostumer string, widUser string) {
	var Session models.Session
	db := database.GetDatabase()
	db.Where("wid_client = ? AND wid_user = ?", widCostumer, widUser).Last(&Session)
	Session.Rate = rate
	db.Save(&Session)
}
