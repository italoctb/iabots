package controllers

import (
	"app/server/adapters"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GPTHandler(c *gin.Context) {
	var requestPayload adapters.ResponseType
	err := c.ShouldBindJSON(&requestPayload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	gptMessages := []Message{}

	gptMessages = append(gptMessages, Message{
		Role:    "user",
		Content: requestPayload.Messages[0].Text.Body,
	})

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

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot send request: " + err.Error(),
		})
		return
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot read response body to bytes: " + err.Error(),
		})
		return
	}
	var aiResponse GPTResponse
	err = json.Unmarshal([]byte(bytes), &aiResponse)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot unmarshall this object: " + err.Error(),
		})
		return
	}
	fmt.Println("xxxxxx")
	fmt.Println(aiResponse.Choices[0].Message.Content)
	fmt.Println("xxxxxx")

	finalResponse := PositusGptMessage{
		To:   requestPayload.Contacts[0].WidSender,
		Type: "text",
		Text: PositusText{Body: aiResponse.Choices[0].Message.Content},
	}

	c.JSON(200, finalResponse)
	Positus := adapters.Positus{}
	err = Positus.SendMessage(requestPayload.Contacts[0].WidSender, aiResponse.Choices[0].Message.Content)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot unmarshall this object: " + err.Error(),
		})
		return
	}

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
	Choices []ChoicesGPT `json:"choices"`
}

type ChoicesGPT struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type PositusGptMessage struct {
	To   string      `json:"to"`
	Type string      `json:"type"`
	Text PositusText `json:"text"`
}

type PositusText struct {
	Body string `json:"body"`
}
