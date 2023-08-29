package controllers

import (
	"app/server/adapters"
	"app/server/bots"
	"app/server/database"
	"app/server/models"
	"app/server/pipelines"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GPTHandler(c *gin.Context) {
	db := database.GetDatabase()

	var requestPayload adapters.ResponseType
	err := c.ShouldBindJSON(&requestPayload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	var Costumer models.Costumer

	Bot := bots.ExampleBot{}
	err = db.First(&Costumer).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot retrieve Costumer: " + err.Error(),
		})

		return
	}
	fmt.Println("Sender: " + requestPayload.Messages[len(requestPayload.Messages)-1].From)
	fmt.Println("Receiver: " + requestPayload.Contacts[len(requestPayload.Contacts)-1].WidReceiver)
	fmt.Println("NewReceiver: " + Costumer.Wid)
	payloadMessage := models.Message{
		WidReceiver: Costumer.Wid,
		WidSender:   requestPayload.Messages[len(requestPayload.Messages)-1].From,
		Message:     requestPayload.Messages[len(requestPayload.Messages)-1].Text.Body,
		ProcessedAt: true,
	}
	err = pipelines.ChainProcessGPT(Bot, Costumer, &payloadMessage)
	if err != nil {
		c.JSON(400, "Erro")
	}
	c.JSON(200, "Sucesso")

}

type GPTPayload struct {
	Model            string       `json:"model"`
	Messages         []MessageGPT `json:"messages"`
	MaxTokens        int          `json:"max_tokens"`
	Temperature      int          `json:"temperature"`
	TopP             int          `json:"top_p"`
	FrequencyPenalty int          `json:"frequency_penalty"`
	PresencePenalty  int          `json:"presence_penalty"`
}

type MessageGPT struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTResponse struct {
	Choices []ChoicesGPT `json:"choices"`
}

type ChoicesGPT struct {
	Index        int        `json:"index"`
	Message      MessageGPT `json:"message"`
	FinishReason string     `json:"finish_reason"`
}

type PositusGptMessage struct {
	To   string      `json:"to"`
	Type string      `json:"type"`
	Text PositusText `json:"text"`
}

type PositusText struct {
	Body string `json:"body"`
}
