package controllers

import (
	"app/server/adapters"
	"app/server/bots"
	"app/server/database"
	"app/server/models"
	"app/server/pipelines"
	"net/http"

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
	payloadMessage := models.Message{
		WidReceiver: Costumer.Wid,
		WidSender:   requestPayload.Messages[len(requestPayload.Messages)-1].From,
		Message:     requestPayload.Messages[len(requestPayload.Messages)-1].Text.Body,
		ProcessedAt: true,
	}
	response, err := pipelines.ChainProcessGPT(Bot, Costumer, &payloadMessage)
	if err != nil {
		c.JSON(400, "Erro GPT Pipeline")
	}
	c.JSON(200, response)

}

func ValidateWebhook(c *gin.Context) {

	verifyToken := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	// Verifique se o token de verificação corresponde ao seu token configurado
	if verifyToken == "dragonballz" {
		// Retorne o valor hub.challenge para concluir a verificação
		c.String(http.StatusOK, challenge)
		return
	}

}

func MetaGPTHandler(c *gin.Context) {
	db := database.GetDatabase()

	var requestPayload adapters.MetaResponseObject
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

	if requestPayload.Entry != nil && requestPayload.Entry[0].Changes != nil && len(requestPayload.Entry[0].Changes[0].Value.Messages) > 0 {

		responseMessages := requestPayload.Entry[0].Changes[0].Value.Messages
		payloadMessage := models.Message{
			WidReceiver: Costumer.Wid,
			WidSender:   responseMessages[0].From,
			Message:     responseMessages[0].Text.Body,
			ProcessedAt: true,
		}
		response, err := pipelines.ChainProcessGPT(Bot, Costumer, &payloadMessage)
		if err != nil {
			c.JSON(400, "Erro GPT Pipeline:"+err.Error())
		}
		c.JSON(200, response)
	} else {
		c.Status(404)
	}

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

type MetaWebhookContest struct {
	VerifyToken string `json:"hub.verify_token"`
	Challenge   string `json:"hub.challenge"`
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
