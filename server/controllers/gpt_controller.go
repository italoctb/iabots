package controllers

import (
	"app/server/adapters"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GPTHandler(c *gin.Context) {
	var requestPayload adapters.ResposeType
	err := c.ShouldBindJSON(&requestPayload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	gptMessages := []Message{}
	for _, message := range requestPayload.Messages {
		gptMessages = append(gptMessages, Message{
			Role:    "user",
			Content: message.Text.Body,
		})
	}

	gptPayload := GPTPayload{
		Model:            "gpt-3.5-turbo",
		Messages:         gptMessages,
		MaxTokens:        1500,
		Temperature:      1,
		TopP:             1,
		FrequencyPenalty: 1,
		PresencePenalty:  1,
	}

	fmt.Printf("%+v", gptPayload)
	gptPayloadBody, err := json.Marshal(gptPayload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot marshal gpt payload: " + err.Error(),
		})
		return
	}

	fmt.Println(string(gptPayloadBody))

	gptPayloadStream := bytes.NewBuffer(gptPayloadBody)
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", gptPayloadStream)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create request: " + err.Error(),
		})
		return
	}

	req.Header.Add("Content-Type", "application/json")
	token := os.Getenv("OPENIA_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	fmt.Printf("%+v", res)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot send request: " + err.Error(),
		})
		return
	}

	defer res.Body.Close()
	body, err := json.Marshal(res.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot read response body: " + err.Error(),
		})
		return
	}
	c.JSON(200, string(body))

}

type GPTPayload struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	MaxTokens        int       `json:"max_tokens"`
	Temperature      int       `json:"temperature"`
	TopP             int       `json:"top_p"`
	FrequencyPenalty int       `json:"frequency_penalty"`
	PresencePenalty  int       `json:"presence_penalty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTResponse struct {
	Choices []Message `json:"choices"`
}
