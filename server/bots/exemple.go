package bots

import (
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

func (l ExampleBot) SendMessage(message string) error {
	db := database.GetDatabase()
	newMessage := models.Message{
		Message:     message,
		ProcessedAt: true}
	db.Create(newMessage)
	return nil
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
	return Session.State
}

func (l ExampleBot) GetFirstTemplate() string {
	db := database.GetDatabase()
	var Template models.Template
	db.First(&Template)
	return strconv.FormatUint(uint64(Template.ID), 10)
}

func (l ExampleBot) GetOptions() []int {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Option").Find(&Template, "ID=?", l.GetState())

	list := []int{}
	pivot := 0
	for _, o := range Template.Options {
		list = append(list, pivot)
		fmt.Println(o)
	}
	return list
}

func (l ExampleBot) GetLink(position int) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Option").Find(&Template, "ID=?", l.GetState())
	return Template.Options[position].Goto
}

func (l ExampleBot) TemplateMessage(state string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Find(&Template, "ID=?", state)
	return Template.GetMessage()
}
