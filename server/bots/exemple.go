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

func (l ExampleBot) FallbackMessage() string {
	return "Opção não existe"
}
func (l ExampleBot) EndMessage() string {
	return "Chegamos ao fim"
}

func (l ExampleBot) SendMessage(message string, receiver string) error {
	db := database.GetDatabase()
	newMessage := models.Message{
		Message:     message,
		WidReceiver: receiver,
		ProcessedAt: true}
	db.Create(&newMessage)
	Positus := adapters.Positus{}
	err := Positus.SendMessage(receiver, message)
	return err
}

func (l ExampleBot) SetState(link string) string {
	db := database.GetDatabase()
	newSession := models.Session{
		State: link,
	}
	db.Create(&newSession)
	return ""
}

func (l ExampleBot) GetState() string {
	var Session models.Session
	db := database.GetDatabase()
	db.Last(&Session)
	if Session.State == "" {
		l.SetState(l.GetFirstTemplate())
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

func (l ExampleBot) GetOptions() []int {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", l.GetState())

	list := []int{}
	pivot := 0
	for _, o := range Template.Options {
		list = append(list, pivot)
		pivot += 1
		fmt.Println(o)
	}
	return list
}

func (l ExampleBot) GetLink(position int) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", l.GetState())
	return Template.Options[position-1].Goto
}

func (l ExampleBot) TemplateMessage(state string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", state)
	return Template.GetMessage()
}
