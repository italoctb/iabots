package models

type ChatGPTConfig struct {
	Model            string  `json:"model"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             float32 `json:"top_p"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
	PresencePenalty  float32 `json:"presence_penalty"`
	Prompt           string  `json:"prompt,omitempty"`
	CustomerID       *uint   `json:"customer_id,omitempty"`
}
