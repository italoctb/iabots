package adapters

import (
	"app/server/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type EmbedingBody struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type EmbedingResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}

func FetchEmbedingFromGPT4(text string) ([]float32, error) {
	gptPayload := EmbedingBody{
		Model: "text-embedding-ada-002",
		Input: text,
	}

	gptPayloadBody, err := json.Marshal(gptPayload)
	if err != nil {
		return nil, err
	}

	gptPayloadStream := bytes.NewBuffer(gptPayloadBody)
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", gptPayloadStream)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	token := os.Getenv("OPENIA_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response EmbedingResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Data[0].Embedding, nil
}

type RoleMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTPayload struct {
	models.ChatGPTConfig
	Messages []RoleMessage `json:"messages"`
}

func GPTChatCompletion(config models.ChatGPTConfig, messages []RoleMessage) (string, error) {
	ChatGPTPayload := ChatGPTPayload{
		ChatGPTConfig: config,
		Messages:      []RoleMessage{},
	}
	ChatGPTPayload.Messages = []RoleMessage{
		{
			Role:    "system",
			Content: ChatGPTPayload.Prompt,
		},
	}
	ChatGPTPayload.Prompt = ""
	ChatGPTPayload.CustomerID = nil
	ChatGPTPayload.Messages = append(ChatGPTPayload.Messages, messages...)

	configBytes, err := json.Marshal(ChatGPTPayload)
	if err != nil {
		return "", err
	}

	configStream := bytes.NewBuffer(configBytes)
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", configStream)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	token := os.Getenv("OPENIA_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := client.Do(req)

	if res.StatusCode > 300 {
		body := make([]byte, res.ContentLength)
		res.Body.Read(body)

		return "", errors.New(string(body))
	}

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var response struct {
		Choices []struct {
			RoleMessage `json:"message"`
		} `json:"choices"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Content, nil
}
