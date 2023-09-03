package pipelines

import (
	"app/server/bots"
	"app/server/database"
	"app/server/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func TemplateResponse(b bots.Bot, c models.Costumer, Message *models.Message) error {
	user := getUserFromMessage(c, *Message)
	state := b.GetStateTemplate(c.Wid, user)
	if state == "end" {
		b.SetState(b.GetFirstTemplate(c.Wid), c.Wid, user)
		return nil
	}
	TemplateMessage := b.TemplateMessage(state)
	err := b.SendMessage(TemplateMessage, c.Wid, Message.WidSender, b.GetSession(c.Wid, user).ID)
	return err
}

func GetRoleMessages(b bots.Bot, c models.Costumer, userNumber string) []RoleMessage {
	db := database.GetDatabase()
	var messages []models.Message
	var roleMessages []RoleMessage
	session := b.GetSession(c.Wid, userNumber)
	fmt.Println("Adicionando contexto de sistema...")
	roleMessages = append(roleMessages, RoleMessage{
		Role:    "system",
		Content: "Você se chama Delillah, uma assistente virtual GPT, muito educada, organizada, e proativa.",
	})
	roleMessages = append(roleMessages, RoleMessage{
		Role:    "system",
		Content: "A primeira interação com o usuário, voce se apresenta",
	})
	fmt.Println("Iniciando busca de mensagens da sessão...")
	db.Where("session_id = ?", session.ID).Find(&messages)
	for _, message := range messages {
		if message.WidSender == c.Wid {

			roleMessages = append(roleMessages, RoleMessage{
				Role:    "assistant",
				Content: message.Message,
			})
			fmt.Println("Assistant: " + message.Message)
		} else if message.WidSender == userNumber {
			roleMessages = append(roleMessages, RoleMessage{
				Role:    "user",
				Content: message.Message,
			})
			fmt.Println("user: " + message.Message)
		}

	}
	return roleMessages
}

func GetGPTResponse(b bots.Bot, c models.Costumer, Message *models.Message) (string, error) {
	user := getUserFromMessage(c, *Message)
	gptMessages := GetRoleMessages(b, c, user)

	gptPayload := GPTPayload{
		Model:            "gpt-3.5-turbo",
		Messages:         gptMessages,
		MaxTokens:        1500,
		Temperature:      0.33,
		TopP:             1,
		FrequencyPenalty: 1,
		PresencePenalty:  1,
	}

	fmt.Printf("%+v", gptPayload)
	gptPayloadBody, err := json.Marshal(gptPayload)
	if err != nil {

		return "", err
	}

	fmt.Println(string(gptPayloadBody))

	gptPayloadStream := bytes.NewBuffer(gptPayloadBody)
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", gptPayloadStream)
	if err != nil {

		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	token := os.Getenv("OPENIA_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := client.Do(req)

	if err != nil {

		return "", err

	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {

		return "", err
	}
	var aiResponse GPTResponse
	err = json.Unmarshal([]byte(bytes), &aiResponse)
	if err != nil {

		return "", err
	}
	for _, message := range aiResponse.Choices {
		fmt.Println("xxxxxx")
		fmt.Println(message.Message.Content)
		fmt.Println("xxxxxx")
	}

	finalResponse := aiResponse.Choices[0].Message.Content

	err = b.SendMessage(finalResponse, c.Wid, Message.WidSender, b.GetSession(c.Wid, user).ID)
	if err != nil {
		return "", err
	}
	return finalResponse, nil
}

func ChangeStateBasedOnSelectedOption(b bots.Bot, c models.Costumer, Message *models.Message) error {
	user := getUserFromMessage(c, *Message)
	Option, err := strconv.Atoi(Message.Message)
	if err != nil && (b.GetStateTemplate(c.Wid, user) != b.GetFirstTemplate(c.Wid)) {
		b.SendMessage(c.FallbackMessage, c.Wid, Message.WidSender, b.GetSession(c.Wid, user).ID)
		return err
	}
	options := b.GetOptions(c.Wid, user)
	if checkStateOptions(b, c, user, options, Option) {
		if Option != 0 {
			b.SendMessage(c.FallbackMessage, c.Wid, Message.WidSender, b.GetSession(c.Wid, user).ID)
		}
		return nil
	} else {
		if strconv.FormatUint(uint64(c.RateTemplateID), 10) == b.GetStateTemplate(c.Wid, user) {
			b.RateSession(Option, c.Wid, user)
			b.SendMessage(c.EndMessage, c.Wid, Message.WidSender, b.GetSession(c.Wid, user).ID)
			b.SetState("end", c.Wid, user)
			return err
		}
	}
	b.SetState(b.GetLink(Option, c.Wid, user), c.Wid, user)
	return err
}

func ChangeStateBasedStatus(b bots.Bot, c models.Costumer, Message *models.Message) (models.Session, error) {
	user := getUserFromMessage(c, *Message)

	// if b.GetSession(c.Wid, user).State == "INITIAL" {
	// 	b.SendMessage("Olá! Me chamo Delillah, sou sua assistente virtual GPT!", c.Wid, user)
	// }

	session := b.SetState("ACTIVE", c.Wid, user)
	return session, nil
}

func getUserFromMessage(c models.Costumer, m models.Message) string {
	if c.Wid == m.WidSender {
		return m.WidReceiver
	}
	return m.WidSender
}

func checkStateOptions(b bots.Bot, c models.Costumer, user string, options []int, Option int) bool {
	if len(options) == 0 {
		return (strconv.FormatUint(uint64(c.RateTemplateID), 10) == b.GetStateTemplate(c.Wid, user) && (Option < 1 || Option > 3))
	}
	return Option > len(options) || Option < 1
}

func ResetStateTemplate(b bots.Bot, c models.Costumer, Message *models.Message) error {
	user := getUserFromMessage(c, *Message)
	var Session models.Session
	db := database.GetDatabase()
	db.Where("wid_costumer = ? AND wid_user = ?", c.Wid, user).Last(&Session)
	if getConditionsToReset(Message.Message, Session.CreatedAt) {
		b.SetState(b.GetFirstTemplate(c.Wid), c.Wid, user)
	}
	return nil
}

func ResetState(b bots.Bot, c models.Costumer, Message *models.Message) (models.Session, error) {
	user := getUserFromMessage(c, *Message)
	session := b.GetSession(c.Wid, user)
	if getConditionsToReset(Message.Message, session.UpdateAt) && session.State != "INITIAL" {
		session = b.SetState("CLOSED", c.Wid, user)
	}
	return session, nil
}

func getConditionsToReset(message string, updatedAt time.Time) bool {
	delayTime := (-30) * time.Minute //(-24) * time.Hour || (-1) * time.Minute
	currentTime := time.Now()
	check := currentTime.Add(delayTime).After(updatedAt) || message == "reset"
	return check
}

func ChainProcess(b bots.Bot, c models.Costumer, Message *models.Message) error {
	ResetStateTemplate(b, c, Message)
	ChangeStateBasedOnSelectedOption(b, c, Message)
	TemplateResponse(b, c, Message)
	Message.ProcessedAt = true
	db := database.GetDatabase()
	db.Save(&Message)
	return nil
}

func ChainProcessGPT(b bots.Bot, c models.Costumer, Message *models.Message) (string, error) {
	session, err := ResetState(b, c, Message)
	if session.State == "CLOSED" {
		Message.SessionID = session.ID
		db := database.GetDatabase()
		db.Create(&Message)

		user := getUserFromMessage(c, *Message)
		b.SendMessage("_Sessão encerrada_", c.Wid, user, session.ID)
		ChangeStateBasedStatus(b, c, Message)
	} else {
		session, err = ChangeStateBasedStatus(b, c, Message)
		Message.SessionID = session.ID
		db := database.GetDatabase()
		db.Create(&Message)
	}

	if err != nil {
		return "", err
	}

	response, err := GetGPTResponse(b, c, Message)
	if err != nil {
		return "", err
	}
	return response, nil
}

type RoleMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTResponse struct {
	Choices []ChoicesGPT `json:"choices"`
}

type ChoicesGPT struct {
	Index        int         `json:"index"`
	Message      RoleMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type PositusGptMessage struct {
	To   string      `json:"to"`
	Type string      `json:"type"`
	Text PositusText `json:"text"`
}

type PositusText struct {
	Body string `json:"body"`
}

type GPTPayload struct {
	Model            string        `json:"model"`
	Messages         []RoleMessage `json:"messages"`
	MaxTokens        int           `json:"max_tokens"`
	Temperature      float32       `json:"temperature"`
	TopP             int           `json:"top_p"`
	FrequencyPenalty int           `json:"frequency_penalty"`
	PresencePenalty  int           `json:"presence_penalty"`
}
