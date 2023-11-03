package bots

import (
	"app/server/adapters"
	"app/server/database"
	"app/server/models"
	"fmt"
	"strconv"
	"time"
)

type ExampleBot struct {
}

func (l ExampleBot) SendMessage(message string, sender string, receiver string, sessionId int) error {
	db := database.GetDatabase()
	newMessage := models.Message{
		Message:     message,
		WidSender:   sender,
		WidReceiver: receiver,
		SessionID:   sessionId,
		ProcessedAt: true}
	db.Create(&newMessage)
	//Positus := adapters.Positus{}
	Meta := adapters.Meta{}
	err := Meta.SendMessage(receiver, message)
	return err
}

func (l ExampleBot) SetState(state string, widCustomer string, widUser string) models.Session {
	session := l.GetSession(widCustomer, widUser)

	db := database.GetDatabase()

	session.State = state
	session.UpdateAt = time.Now()
	db.Save(&session)
	return session
}

func (l ExampleBot) GetStateTemplate(widCustomer string, widUser string) string {
	var Session models.Session
	db := database.GetDatabase()
	fmt.Print("wid customer = " + widCustomer)
	db.Where("wid_customer = ? AND wid_user = ?", widCustomer, widUser).Last(&Session)
	fmt.Print("Sessão: " + Session.State)
	if Session.State == "" || Session.State == "0" {
		l.SetState(l.GetFirstTemplate(widCustomer), widCustomer, widUser)
		return l.GetFirstTemplate(widCustomer)
	}
	return Session.State
}

func (l ExampleBot) GetSession(widCustomer string, widUser string) models.Session {
	var session models.Session
	db := database.GetDatabase()
	fmt.Print("wid customer = " + widCustomer)
	err := db.Where("wid_customer = ? AND wid_user = ? AND state != ?", widCustomer, widUser, "CLOSED").Last(&session).Error
	if err != nil {
		newSession := models.Session{
			State:       "INITIAL",
			WidCustomer: widCustomer,
			WidUser:     widUser,
		}
		db.Create(&newSession)
		fmt.Println("NOVA SESSÃO!")
		PrintSession(newSession)
		return newSession
	}
	fmt.Println("SESSÃO ENCONTRADA!")
	PrintSession(session)
	return session
}

func PrintSession(session models.Session) {
	fmt.Println("Sessão: ")
	fmt.Println("id: " + strconv.Itoa(session.ID))
	fmt.Println("state: " + session.State)
	fmt.Println("widUser: " + session.WidUser)
	fmt.Println("widCustomer: " + session.WidCustomer)
	fmt.Println("createdAt: " + session.CreatedAt.String())
	fmt.Println("updatedAt: " + session.UpdateAt.String())
	fmt.Println()
}

func (l ExampleBot) GetFirstTemplate(widCustomer string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "wid_customer = ? AND is_first=?", widCustomer, true)
	return strconv.FormatUint(uint64(Template.ID), 10)
}

func (l ExampleBot) GetOptions(widCustomer string, widUser string) []int {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", l.GetStateTemplate(widCustomer, widUser))

	list := []int{}
	pivot := 0
	for _, o := range Template.Options {
		list = append(list, pivot)
		pivot += 1
		fmt.Println(o)
	}
	return list
}

func (l ExampleBot) GetLink(position int, widCustomer string, widUser string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", l.GetStateTemplate(widCustomer, widUser))
	return Template.Options[position-1].Goto
}

func (l ExampleBot) TemplateMessage(state string) string {
	db := database.GetDatabase()
	var Template models.Template
	db.Preload("Options").Find(&Template, "ID=?", state)
	return Template.GetMessage()
}

func (l ExampleBot) RateSession(rate int, widCustomer string, widUser string) {
	var Session models.Session
	db := database.GetDatabase()
	db.Where("wid_customer = ? AND wid_user = ?", widCustomer, widUser).Last(&Session)
	Session.Rate = rate
	db.Save(&Session)
}
